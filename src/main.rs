mod proverb;

fn main() {
    let words =  vec!["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"];
    println!("{}", proverb::proverb(words));
}
