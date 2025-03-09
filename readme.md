Problem: Concurrent Web Scraper
Task
Write a Go program that concurrently scrapes the titles of multiple web pages using goroutines. The program should:

Accept a list of URLs.
Launch a goroutine for each URL to fetch the page and extract the title (the text inside the <title> tag).
Collect the titles in a shared slice, ensuring thread safety with a sync.Mutex.
Use a sync.WaitGroup to wait for all goroutines to complete.
Print the list of titles once all pages have been scraped.
Requirements
Concurrency: Use goroutines to fetch and process each URL concurrently.
Synchronization: Protect the shared slice of titles with a sync.Mutex to avoid race conditions.
Completion: Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish.
Error Handling: Gracefully handle errors (e.g., invalid URLs, network issues, or missing <title> tags) by including error messages in the output.
Example Input
Use the following URLs to test your program:

"https://golang.org"
"https://example.com"
"https://nonexistent-site.xyz" (for error handling)
Expected Output
text

Collapse

Wrap

Copy
Titles:
- The Go Programming Language
- Example Domain
- Error scraping https://nonexistent-site.xyz: Get "https://nonexistent-site.xyz": dial t