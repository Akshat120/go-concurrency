# Problem: Concurrent Web Scraper
**Task: Write a Go program that concurrently scrapes the titles of multiple web pages using goroutines. The program should:** 
- Accept a list of URLs.
- Launch a `goroutine` for each URL to fetch the page and extract the title (the text inside the <title> tag).
- Collect the titles in a shared slice, ensuring thread safety with a `sync.Mutex`.
- Use a `sync.WaitGroup` to wait for all goroutines to complete.
- Print the list of titles once all pages have been scraped.

## Requirements
- **Concurrency:** Use goroutines to fetch and process each URL concurrently.
- **Synchronization:** Protect the shared slice of titles with a sync.Mutex to avoid race conditions.
- **Completion:** Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish.
- **Error Handling:** Gracefully handle errors (e.g., invalid URLs, network issues, or missing <title> tags) by including error messages in the output.

### Example Input
Use the following URLs to test your program:
- "https://golang.org"
- "https://example.com"
- "https://nonexistent-site.xyz" (for error handling)
Expected Output
```
Titles:
- The Go Programming Language
- Example Domain
- Error scraping https://nonexistent-site.xyz: Get "https://nonexistent-site.xyz": dial tcp: lookup nonexistent-site.xyz: no such host
```
### Output - 
<img width="968" alt="Screenshot 2025-03-09 at 9 32 52 PM" src="https://github.com/user-attachments/assets/652f5ee8-21a0-4f5a-8fa6-b719fd44774c" />


