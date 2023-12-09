use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;

use regex::Regex;

fn main() {
    let args: Vec<String> = env::args().collect();

    // File hosts.txt must exist in the current path
    if let Ok(lines) = read_lines(&args[1]) {
        // Consumes the iterator, returns an (Optional) String
        let mut tot1 = 0;
        let mut tot2 = 0;

        for line in lines {
            if let Ok(line) = line {
                println!("------------------");
                println!("line : {}", line);
                let n1 = extract_first_and_last_numbers(&line);
                println!("n1={}", n1);
                tot1 = tot1 + n1;

                let l2 = substitute_literal_numbers(&line);
                println!("after literal decoding : {}", l2);
                let n2 = extract_first_and_last_numbers(&l2);
                println!("n2={}", n2);
                tot2 = tot2 + n2;
            }
        }
        println!("Total 1 : {}", tot1);
        println!("Total 2 : {}", tot2);
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn extract_first_and_last_numbers(input: &str) -> i32 {
    // Define a regex pattern to match numbers
    let re = Regex::new(r"[0-9]").unwrap();

    // Collect all numbers from the input string
    let numbers: Vec<i32> = re
        .find_iter(input)
        .filter_map(|m| m.as_str().parse().ok())
        .collect();

    println!("Numbers : {} - {:?}", numbers.len(), numbers);
    // Check if there are at least two numbers
    if numbers.len() >= 2 {
        // Extract the first and last numbers
        let first_number = numbers[0];
        let last_number = *numbers.last().unwrap();
        first_number * 10 +  last_number
    } else if numbers.len() == 1 {
        let first_number = numbers[0];
        first_number * 10 + first_number
    } else {
        0
    }
}

fn reverse(s: &str) -> String {
    s.chars().rev().collect()
}

fn convert(s: &str) -> &str {
    let n = match s {
        "one" => "1",
        "two" => "2",
        "three" => "3",
        "four" => "4",
        "five" => "5",
        "six" => "6",
        "seven" => "7",
        "eight" => "8",
        "nine" => "9",
        &_ => "*"
    };
    n
}

fn substitute_literal_numbers(input: &str) -> String {
    let result = input.to_string();
    let mut out = String::new();

    let re = Regex::new(r"one|two|three|four|five|six|seven|eight|nine|[0-9]").unwrap();
    let re_rev = Regex::new(r"eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9]").unwrap();
    let num_re = Regex::new(r"[0-9]").unwrap();

    let s = re.find(&result).unwrap().as_str();
    if num_re.is_match(s) {
        out.push_str(s);
    } else {
        out.push_str(convert(s));
    }

    // now reverse
    let reversed = reverse(&result);
    let e = re_rev.find(reversed.as_str()).unwrap().as_str();
    if num_re.is_match(e) {
        out.push_str(e);
    } else {
        let r_match = reverse(e);
        println!("rev {}", r_match);
        out.push_str(convert(r_match.as_str()));
    }

    out
}