#![feature(async_closure)]

#[macro_use]
extern crate serde_derive;

use reqwest;

#[derive(Serialize, Deserialize, Debug)]
struct Story {
    by: String,
    descendants: i64,
    id: i64,
    //kids : Vec<i64>,
    score: i64,
    time: i64,
    title: String,
    r#type: String,
    url: String,
}

async fn async_cloc(i: i32, story_list: &mut Vec<Story>) -> Result<(), Box<dyn std::error::Error>> {
    let j: Vec<i32> = vec![1, 2, 3];
    let client = reqwest::Client::builder().use_sys_proxy().build()?;
    // 闭包内用不了 try
    // https://www.reddit.com/r/rust/comments/58yk8y/questionmark_or_try_within_map_closure/
    let json = client
        .get(
            format!(
                "https://hacker-news.firebaseio.com/v0/item/{}.json",
                j[i as usize]
            )
            .as_str(),
        )
        .send()
        .await?
        .text()
        .await?;
    let story: Story = serde_json::from_str(json.as_str()).unwrap();
    story_list.push(story);
    Ok(())
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut story_list = vec![];

    /*
    这里怎么都编译不过
    (0..2).for_each(async move |i| {
        async_cloc(i, &mut story_list).await.unwrap();
    });
    */


    for i in (0..2) {
        async_cloc(i, &mut story_list).await;
    }
    println!("{:?}", story_list);
    Ok(())
}
