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

        let mut ring = Vec::new();

        // lines loop
        for line in lines {
            if let Ok(line) = line {
                println!("------------------");
                println!("line : {}", line);

                // provide a "dummy" empty line -1
                if ring.len() == 0 {
                    ring.push(".".repeat(line.as_str().len()));
                } else if ring.len() == 3 {
                    // or remove out of scope lines
                    ring.remove(0);
                }
                // add line    
                ring.push(line);
                
                // ensure ring full
                if ring.len() < 3 {
                    continue;
                }
                
                print_ring(&ring);
                // check on 2nd line
                tot1 = tot1 + check_line(&ring).iter().sum::<u32>();

                // check gear
                tot2 = tot2 + check_gear(&ring).iter().sum::<u32>();
            }
        }

        // process last line !
        // add a pad line
        ring.remove(0);
        ring.push(".".repeat(ring[0].len()));

        print_ring(&ring);
        // check on 2nd line
        tot1 = tot1 + check_line(&ring).iter().sum::<u32>();

        // check gear
        tot2 = tot2 + check_gear(&ring).iter().sum::<u32>();

        println!("tot1 : {}", tot1);
        println!("tot2 : {}", tot2);
    }      
}

fn print_ring(ring: &Vec<String>) {
    println!("------------------");
    for l in ring {
        println!("{}", l);
    }
    println!("------------------");
}

fn check_gear(ring : &Vec<String>) -> Vec<u32> {
    let re_gear = Regex::new(r"\*").unwrap();
    let re = Regex::new(r"[0-9]+").unwrap();

    let len = ring[0].len();

    let mut ret = Vec::new();

    // locate gears
    for g in re_gear.find_iter(&ring[1]) {
        let gear_footprint = calc_footprint(len, g.start(), g.end());
        let gs = gear_footprint.0;
        let ge = gear_footprint.1-1;

        println!("gear {}-{} {:?}",gs, ge, g);
        let mut pair = Vec::new();
        for rl in ring {
            for nm in re.find_iter(&rl) {
                let cs =  nm.start();
                let ce = nm.end()-1;
                if gs <= ce && cs <= ge {
                    let nm_val = nm.as_str().parse().unwrap();
                    pair.push(nm_val);
                }
            }
        }
        if pair.len() == 2 {
            println!("gear pair {:?}", pair);
            ret.push(pair.iter().product());
        }
    }

    ret
}
 
fn check_line(ring: &Vec<String>) -> Vec<u32> {
    let re = Regex::new(r"[0-9]+").unwrap();
    let mut ret = Vec::new();
    for n in re.find_iter(&ring[1]) {
        println!("# {:?}", n);
        let len = ring[1].len();
        let start = n.start();
        let end = n.end();
        let chars: Vec<char> = ring[1].chars().collect();

        // check left, right, top, bottom
        let valid = (start > 0 && chars[start-1] != '.')  || 
            (end < len && chars[end] != '.') ||
            check_near(ring, 0, start, end) || 
            check_near(ring, 2, start, end) 
            ;
   
        let n_val = n.as_str().parse().unwrap();
        println!("valid {} {}", valid, n_val);
        if valid {
            ret.push(n_val);
        }
    }
    ret
}

fn calc_footprint(len : usize, start: usize, end: usize) -> (usize, usize) {

    let cs = if start > 0 { 
            start - 1 
        } else { 
            0 
        };
    let ce = if end < len  { 
        end + 1
    } else {
        end 
    };
    (cs, ce)
}

fn check_near(ring : &Vec<String>, idx : usize, start: usize, end: usize) -> bool {
    let len = ring[idx].len();
    let footprint = calc_footprint(len, start, end);
    let cs = footprint.0;
    let ce = footprint.1;
    
    let slice = ring[idx].get(cs as usize..ce as usize);
    let mask = ".".repeat(ce-cs);
    let valid = slice.unwrap() != mask.as_str();
    println!("slice[{}] {:?} {}-{}", idx, slice, cs, ce);
    
    valid
}


// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
}

