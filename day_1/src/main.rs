use std::fs::File;
use std::io::{prelude::*, BufReader};

fn read_from_file(path: &str) -> Vec<i32> {
    let file = File::open(path).unwrap();
    let reader = BufReader::new(file);

    let mut output = Vec::new();
    let mut count = 0;
    for line in reader.lines() {
        if let Ok(text) = line {
            if let Ok(num) = text.parse::<i32>() {
                count += num;
            } else {
                output.push(count);
                count = 0;
            }
        } else {
            panic!("failed to parse")
        }
    }
    output.push(count);
    output
}

fn main() {
    let mut calorie_count = read_from_file("input/input.txt");
    calorie_count.sort_by(|a, b| b.cmp(a));

    println!("part 1: {:?}", calorie_count[0]);

    println!(
        "part 2: {:?}",
        calorie_count[0] + calorie_count[1] + calorie_count[2]
    );
}
