use reqwest::Client;
use reqwest::StatusCode;

fn main() -> Result<(), Box<::std::error::Error>> {
    let client = Client::new();

    let mut resp = client
        .get("https://www.collinsdictionary.com/autocomplete/?dictCode=english&q=ab")
        .send()?;

    /*
    match resp.status() {
        StatusCode::OK => println!("success!"),
        StatusCode::PAYLOAD_TOO_LARGE => {
            println!("Request payload is too large!");
        }
        s => println!("Received response status: {:?}", s),
    };
    */
    /*
    let js: serde_json::Value = resp.json()?;
    println!("{}", js);
    */
    //println!("{}", resp.json()?);
    println!("{}", resp.json::<serde_json::Value>()?);
    Ok(())
}
