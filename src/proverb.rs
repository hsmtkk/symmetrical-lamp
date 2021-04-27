pub fn proverb(words:Vec<&str>) -> String {
    let mut phrases = Vec::new();
    for i in 0..words.len()-1 {
        phrases.push(format!("For want of a {} the {} was lost.", words[i], words[i+1]));
    }
    phrases.push("And all for the want of a nail.".to_string());
    return phrases.join("\n");   
}