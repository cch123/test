use mio::event::Event;
use mio::net::{TcpListener, TcpStream};
use mio::{Events, Interest, Poll, Registry, Token};
use std::io::{self, Read, Write};
use std::collections::HashMap;
use std::str::from_utf8;

const SERVER: Token = Token(0);

const DATA: &[u8] = b"hello world\n";

fn main() -> io::Result<()> {
    let mut poll = Poll::new()?;
    let mut events = Events::with_capacity(128);

    let addr = "127.0.0.1:9999".parse().unwrap();

    let mut server = TcpListener::bind(addr)?;

    poll.registry()
        .register(&mut server, SERVER, Interest::READABLE)?;

    let mut connections = HashMap::new();

    // Unique token for each incoming connection.
    let mut unique_token = Token(SERVER.0 + 1);

    loop {
        poll.poll(&mut events, None)?;

        for event in events.iter() {
            match event.token() {
                SERVER => {
                    // Received an event for the TCP server socket.
                    // Accept an connection.
                    let (mut connection, addr) = server.accept()?;
                    println!("accept a conn : {}", addr);

                    let token = next(&mut unique_token);
                    poll.registry().register(
                        &mut connection, token, Interest::READABLE.add(Interest::WRITABLE),
                    );
                    connections.insert(token, connection);
                }
                token => {
                    let done = if let Some(connection) = connections.get_mut(&token) {
                        handle_connection_event(poll.registry(), connection, event)?
                    } else {
                        false
                    };
                    if done {
                        connections.remove(&token);
                    }
                }
            }
        }
    }
}

fn next(current: &mut Token) -> Token {
    let next = current.0;
    current.0 += 1;
    Token(next)
}

fn handle_connection_event(
    registry: &Registry,
    connection: &mut TcpStream,
    event: &Event,
) -> io::Result<bool> {
    if event.is_writable() {
        match connection.write(DATA) {
            Ok(n) if n < DATA.len() => {
                return Err(io::ErrorKind::WriteZero.into());
            }
            Ok(_) => {
                registry.reregister(connection, event.token(), Interest::READABLE)?
            }
            Err(ref err) if would_block(err) => {}
            // Got interrupted (how rude!), we'll try again.
            Err(ref err) if interrupted(err) => {
                return handle_connection_event(registry, connection, event)
            }
            // Other errors we'll consider fatal.
            Err(err) => return Err(err),
        }
    }

    if event.is_readable() {
        let mut connection_closed = false;
        let mut received_data = Vec::with_capacity(4096);

        loop {
            let mut buf = [0;256];
            match connection.read(&mut buf) {
                Ok(0) => {
                    connection_closed = true;
                    break;
                }
                Ok(n) => {
                    received_data.extend_from_slice(&buf[..n]);
                }
                Err(ref err) if would_block(err) => break,
                Err(ref err) if interrupted(err) => continue,
                // Other errors we'll consider fatal.
                Err(err) => return Err(err),
            }
        }

        if let Ok(str_buf) = from_utf8(&received_data) {
            println!("Received data: {}", str_buf.trim_end());
        } else {
            println!("Received (none UTF-8) data: {:?}", &received_data);
        }

        if connection_closed {
            println!("Connection closed");
            return Ok(true);
        }
    }

    Ok(false)
}

fn would_block(err: &io::Error) -> bool {
    err.kind() == io::ErrorKind::WouldBlock
}

fn interrupted(err: &io::Error) -> bool {
    err.kind() == io::ErrorKind::Interrupted
}