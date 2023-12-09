use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;
use std::collections::HashMap;

// use regex::Regex;

fn main() {
    let args: Vec<String> = env::args().collect();  

    let backpack = HashMap::from([
        ("red", 12),
        ("green", 13),
        ("blue", 14),
    ]);

   // File hosts.txt must exist in the current path
   if let Ok(lines) = read_lines(&args[1]) {
        // Consumes the iterator, returns an (Optional) String

        let mut tot1 = 0;
        let mut tot2 = 0;

        for line in lines {
            // hold game state
            let mut g_state = HashMap::new();
            let mut p_state = HashMap::new();

            if let Ok(line) = line {
                println!("------------------");
                println!("{}", line);
                let mut game = 0;
                
                let mut valid = true;
                for r in line.split([';',':']) {
                    if r.starts_with("Game ") {
                        game = r.strip_prefix("Game ").unwrap().parse().unwrap();
                        println!("Game set to {}", game);
                        p_state.clear();
                    } else {
                        println!("round [{}] ",r);
                        g_state.clear();
                        for p in r.split(",") {
                            let e: Vec<&str> = p.trim().split(' ').collect();
                            let q: u32  = e[0].parse().unwrap();
                            let c = e[1];

                            // update game state if still valid
                            if valid {
                                if !g_state.contains_key(c) {
                                    g_state.insert(c, q);
                                } else {
                                    g_state.insert(c, g_state.get(c).unwrap() + q);
                                }
                                valid = check_state(&g_state, &backpack);
                            }

                            // update power state
                            if !p_state.contains_key(c) {
                                p_state.insert(c, q);
                            } else if p_state.get(c).unwrap() < &q {
                                p_state.insert(c, q);
                            }
                        }
                    }
                }
  
                println!("");
  
                if valid {
                    tot1 = tot1 + game;
                }

                let mut pv = 1;
                for p in p_state.values() {
                    pv = pv * p;
                }
                tot2 = tot2 + pv;

                println!("{} game state {} : {:?}",game,  valid, g_state);
                println!("{} game power {} : {:?}",game, pv, p_state);
            }
        }

        println!("tot1 : {}", tot1);
        println!("tot2 : {}", tot2);
    }   
}

fn check_state(state: &HashMap<&str,u32>, backpack : &HashMap<&str,u32>) -> bool {
    println!("game state {:?}", state);
    for (c, q) in state.iter() {
        if q > backpack.get(c).unwrap() {
            return false;
        }
    }
    true
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
