package main

import "fmt"

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	return fmt.Sprintf("%s.%s", locale, domain)

}

func main() {
	fmt.Println(DomainForLocale("code-basics.com", ""))
	fmt.Println(DomainForLocale("code-basics.com", "ru"))
	fmt.Println(DomainForLocale("code-basics.com", "vn"))
}
