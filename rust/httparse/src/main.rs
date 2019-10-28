use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    let mut headers = [httparse::EMPTY_HEADER; 16];
    let mut req = httparse::Request::new(&mut headers);

    let buf = b"GET /index.html HTTP/1.1\r\nHost";
    assert!(req.parse(buf)?.is_partial());
    println!("{:?}", req);

    // a partial request, so we try again once we have more data

    let buf = b"GET /index.html HTTP/1.1\r\nHost: example.domain\r\n\r\n";
    assert!(req.parse(buf)?.is_complete());
    println!("{:?}", req);
    Ok(())
}
