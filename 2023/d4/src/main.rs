use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;
use std::collections::HashMap;

// use regex::Regex;

const TWO: u32 = 2;

fn main() {
    let args: Vec<String> = env::args().collect();

    // File hosts.txt must exist in the current path
    if let Ok(lines) = read_lines(&args[1]) {
        // Consumes the iterator, returns an (Optional) String
        let mut tot1 = 0;
        let mut tot2 = 0;

        let mut  won_cards: HashMap<String, u32> = HashMap::new();

         // lines loop
        for line in lines {
            if let Ok(line) = line {
                println!("------------------");
                println!("line : {}", line);

                let s1: Vec<&str> = line.split(':').collect();
                let s2: Vec<&str>  = s1[1].split('|').collect();

                println!("s1 :{:?} s2 :{:?}", s1, s2);

                let card: u32 = s1[0].strip_prefix("Card ").unwrap().trim().parse().unwrap();

                let winning: Vec<u32> = s2[0].trim().split_ascii_whitespace().map(|x: &str| x.trim().parse::<u32>().unwrap()).collect();
                let num: Vec<u32> = s2[1].trim().split_ascii_whitespace().map(|x: &str| x.trim().parse::<u32>().unwrap()).collect();
                
                println!("winning :{:?} num :{:?}", winning, num);
            
                let wins: Vec<u32> = num.into_iter().filter(|x| winning.contains(x)).collect();
                
                let win_count = wins.len(); // wins.into_iter().count::<u32>() 
                if win_count  > 0 {
                    tot1 = tot1 + TWO.pow(win_count as u32 - 1);
                    println!("points : {}", TWO.pow(win_count as u32 - 1));

                    let cur = card.to_string();
                    if !won_cards.contains_key(&cur) {
                        won_cards.insert(cur.clone(), 1);
                    } else {   
                        won_cards.insert(cur.clone(), won_cards.get(&cur).unwrap() + 1);
                    }
                    let cur_qty = won_cards.get(&cur).unwrap().clone();
                    println!("cur_qty {}", cur_qty);

                    // add winning next card counts
                    for c in card + 1 .. card + 1 + win_count as u32{
                        let key: String = c.to_string();
                        if !won_cards.contains_key(&key) {
                            won_cards.insert(key, 1 * cur_qty);
                        } else {
                            let nv = won_cards.get(&key).unwrap() +  1 * cur_qty;
                            won_cards.insert(key, nv);
                        }
                    }
                } else {
                    // count also unwon cards
                    tot2 = tot2 + 1;
                }
                println!("Card : {} {:?}", card, wins);
                println!("Won cards : {:?}", won_cards);
             
            }
        }

        tot2 = tot2 + won_cards.values().sum::<u32>();
        println!("tot1 : {}", tot1);
        println!("tot2 : {}", tot2);
    }
}


// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
}

