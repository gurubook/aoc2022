use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;

// use regex::Regex;

fn main() {
    let args: Vec<String> = env::args().collect();

    // File hosts.txt must exist in the current path
    if let Ok(lines) = read_lines(&args[1]) {
        // Consumes the iterator, returns an (Optional) String
        let mut tot1 = 0;
        let mut tot2 = 0;

         // lines loop
        for line in lines {
            if let Ok(line) = line {
                println!("------------------");
                println!("line : {}", line);
                                
            }
        }

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

