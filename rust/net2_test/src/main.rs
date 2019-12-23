use net2::unix::UnixTcpBuilderExt;
use net2::TcpBuilder;
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    let v = TcpBuilder::new_v4().unwrap();
    let vv = v.reuse_port(true).unwrap().bind("0.0.0.0:12345").unwrap().listen(128).unwrap();
    let l = TcpListener::from_std(vv);
    dbg!(&l);
}

