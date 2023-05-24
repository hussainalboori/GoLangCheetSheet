Certainly! Here's a cheat sheet for the `strings` module in Go:

**Package Import:**
```go
import "strings"
```

**Common Functions:**

1. **`strings.Contains(s, substr string) bool`:** Checks if the string `s` contains the substring `substr`. Returns `true` if `substr` is found, `false` otherwise.

2. **`strings.HasPrefix(s, prefix string) bool`:** Checks if the string `s` starts with the specified `prefix`. Returns `true` if `s` has the prefix, `false` otherwise.

3. **`strings.HasSuffix(s, suffix string) bool`:** Checks if the string `s` ends with the specified `suffix`. Returns `true` if `s` has the suffix, `false` otherwise.

4. **`strings.Index(s, substr string) int`:** Returns the index of the first occurrence of `substr` in the string `s`. If `substr` is not found, it returns -1.

5. **`strings.LastIndex(s, substr string) int`:** Returns the index of the last occurrence of `substr` in the string `s`. If `substr` is not found, it returns -1.

6. **`strings.Replace(s, old, new string, n int) string`:** Replaces `n` occurrences of `old` with `new` in the string `s`. If `n` is negative, all occurrences are replaced.

7. **`strings.Split(s, sep string) []string`:** Splits the string `s` into a slice of substrings using the separator `sep`. Returns a slice of strings.

8. **`strings.Join(a []string, sep string) string`:** Concatenates the elements of the string slice `a` into a single string, using the separator `sep` between each element.

9. **`strings.ToLower(s string) string`:** Converts all characters in the string `s` to lowercase and returns the resulting string.

10. **`strings.ToUpper(s string) string`:** Converts all characters in the string `s` to uppercase and returns the resulting string.

11. **`strings.TrimSpace(s string) string`:** Returns a copy of the string `s` with leading and trailing white space removed.

12. **`strings.Trim(s string, cutset string) string`:** Returns a slice of the string `s` with all leading and trailing characters contained in `cutset` removed.

These are some of the commonly used functions from the `strings` module in Go. You can find more functions and their details in the official Go documentation: https://golang.org/pkg/strings/

Hope this helps you with your Go string manipulations!
