/** 
  use net2::TcpBuilder;

let tcp = TcpBuilder::new_v4().unwrap();
tcp.reuse_address(true).unwrap()
   .only_v6(false).unwrap();

let mut stream = tcp.connect("127.0.0.1:80").unwrap();
*/

use net2::unix::UnixTcpBuilderExt;
use net2::TcpBuilder;
use tokio::net::TcpListener;
use tokio::net::TcpStream;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>>{
    let tcp = TcpBuilder::new_v4().unwrap();
    let vv = tcp.reuse_port(true).unwrap().bind("0.0.0.0:12345").unwrap().listen(128).unwrap();
    let l = TcpListener::from_std(vv);
    dbg!(&l);

    client_reuseport().await?;
    Ok(())
}

async fn client_reuseport() -> Result<(), Box<dyn std::error::Error>>{
    let tcp = TcpBuilder::new_v4().unwrap();
    tcp.reuse_port(true).unwrap().bind("0.0.0.0:12345").unwrap();
    let stream = TcpStream::from_std(tcp.connect("220.181.38.148:80").unwrap())?;

    let tcp = TcpBuilder::new_v4().unwrap();
    tcp.reuse_port(true).unwrap().bind("0.0.0.0:12345").unwrap();
    let stream2 = TcpStream::from_std(tcp.connect("111.13.149.108:80").unwrap())?;

    dbg!(&stream, stream2);
    Ok(())
}

