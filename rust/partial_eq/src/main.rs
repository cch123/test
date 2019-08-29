#[derive(Debug, PartialEq, Clone)]
pub enum ListItemType {
    OrderedItem,
    UnorderedItem,
    LabeledItem,
}

#[derive(Debug, Clone)]
pub struct ListItem {
    pub typ: ListItemType,
    pub level: i8,
    pub children: Vec<ListItem>,
    pub content: String,
}

impl PartialEq for ListItem {
    fn eq(&self, other: &Self) -> bool {
        return other.typ == self.typ && other.level == self.level;
    }
}

fn main() {
    let mut a = ListItem{
        typ : ListItemType::OrderedItem,
        level: 2,
        children:vec![],
        content : "abc".to_string(),
    };
    let mut b = ListItem{
        typ : ListItemType::OrderedItem,
        level: 2,
        children:vec![],
        content : "abc".to_string(),
    };
    let mut x = vec![];
    x.push(a);
    println!("{}", x.contains(&b));
}
