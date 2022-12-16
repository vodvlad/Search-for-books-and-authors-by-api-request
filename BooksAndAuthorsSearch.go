package main

import (
	"bufio"
	"os"
	"io/ioutil"
    "net/http"
    "fmt"
	"strings"
	"encoding/json"
	"sort"

)

type ResponseAuthor struct {
    NumFound       int `json:"numFound"`
    Docs       []struct {
        Name     string `json:"name"`
        BirthDate string `json:"birth_date"`
        AlternateNames []string `json:"alternate_names"`
    } `json:"docs"`
}

type ResponseAuthorBooks struct {
    NumFound       int `json:"numFound"`
    Docs       []struct {
        Title     string `json:"title"`
        FirstPublishYear int `json:"first_publish_year"`

    } `json:"docs"`
}

type ResponseBook struct {
    NumFound       int `json:"numFound"`
    Docs       []struct {
        Title     string `json:"title"`
        FirstPublishYear int `json:"first_publish_year"`
        NumberOfPagesMedian int `json:"number_of_pages_median"`
        BookPublisher []string `json:"publisher"`
        PublishPlace []string `json:"publish_place"`
        BookLanguage []string `json:"language"`
		Authors []string `json:"author_name"`
    } `json:"docs"`
}


func GetDataAuthor(link string) ResponseAuthor{
    resp, err := http.Get(link)
    if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var result ResponseAuthor
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }
    return result
}
func GetDataBooksAuthor(link string) ResponseAuthorBooks{
    resp, err := http.Get(link)
    if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var result ResponseAuthorBooks
    if err := json.Unmarshal(body, &result); err != nil {  
        fmt.Println("Can not unmarshal JSON")
    }
    return result
}

func GetDataBook(link string) ResponseBook{
    resp, err := http.Get(link)
    if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var result ResponseBook
    if err := json.Unmarshal(body, &result); err != nil { 
        fmt.Println("Can not unmarshal JSON")
    }
    return result
}



func readString(symbol string) string{
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.ReplaceAll(name, " ", symbol)
	name = strings.ReplaceAll(name, "\n", "")
	name = strings.ToLower(name)
	return name
}


func main() {

	printText("\nYou are welcome!\n \nThis program is designed to search for books, authors, collections and much more.\n \nPlease select an action:\n \n[1] Search by Author Name\n \n[2] Search by book title\n \n[3] GoLang tasks\n")
	fmt.Printf("Please enter a number: ")

	var type_search string
	fmt.Scanf("%s", &type_search)

	fmt.Println("")
	if type_search == "1" {
		printText("\nPlease enter the author's first and last name.\n \nExample: Dan Brown\n")
		fmt.Print("Enter a name: ")
		name := readString("%20")

		result := GetDataAuthor("http://openlibrary.org/search/authors.json?q=" + name)
	
		var AlternateNamesList []string
	
		fmt.Println("Alternative names:", result.Docs[0].Name)
		for i := 0; i < len(result.Docs[0].AlternateNames); i++{
			AlternateNamesList = append(AlternateNamesList, result.Docs[0].AlternateNames[i])
		}
		fmt.Println("Alternative names: ")
		for i := 0; i < len(AlternateNamesList); i++{
			fmt.Println("   ",AlternateNamesList[i])
		}
		
		fmt.Println("Date of birth:",result.Docs[0].BirthDate)
	
		resultBook := GetDataBooksAuthor("http://openlibrary.org/search.json?author=" + name)
	
		var books []string
		var firstPublishYr []int
		for _, rec := range resultBook.Docs {
				books = append(books, rec.Title)
				firstPublishYr = append(firstPublishYr, rec.FirstPublishYear)
		}
		fmt.Println("Books:")
		for i := 0; i < len(books); i++{
			fmt.Println("   ", books[i], firstPublishYr[i])
		}

		fmt.Println("+----------------------------------------------------+")



	} else if type_search == "2"{
		printText("\nplease write the title of the book.\n \nExample: Mastering C++\n")
		fmt.Print("Enter the title: ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		name = strings.ReplaceAll(name, "\n", "")
		origName := name
		if len(origName) > 25{
			fmt.Println("Incorrect title")
			os.Exit(0)
		}
		name = strings.ReplaceAll(name, " ", "%20")
		name = strings.ToLower(name)
		result := GetDataBook("http://openlibrary.org/search.json?title=" + name)
	
		q := 0
		for i := 0; i <= 98; i++{
			if(i == 98){
				q = 0
				break
			}
			if result.Docs[i].Title == origName{
				break
			} else {
				q = q + 1
			}
		}
		fmt.Println("Book:", result.Docs[q].Title)
		fmt.Println("First publish year:", result.Docs[q].FirstPublishYear)
		fmt.Println("Number of pages:", result.Docs[q].NumberOfPagesMedian)
		fmt.Println("Authors: ")
		var AuthorsList []string
		for i := 0; i < len(result.Docs[q].Authors); i++{
			AuthorsList = append(AuthorsList, result.Docs[q].Authors[i])
		}
		for i := 0; i < len(AuthorsList); i++{
			fmt.Println("       ",AuthorsList[i])
		}
		fmt.Println("Publishers: ")
		var PublishersList []string
		for i := 0; i < len(result.Docs[q].BookPublisher); i++{
			PublishersList = append(PublishersList, result.Docs[q].BookPublisher[i])
		}
		for i := 0; i < len(PublishersList); i++{
			fmt.Println("       ",PublishersList[i])
		}
		fmt.Println("Publish place: ")
		var PublishPlaceList []string
		for i := 0; i < len(result.Docs[q].PublishPlace); i++{
			PublishPlaceList = append(PublishPlaceList, result.Docs[q].PublishPlace[i])
		}
		for i := 0; i < len(PublishPlaceList); i++{
			fmt.Println("       ",PublishPlaceList[i])
		}
		fmt.Println("Language: ")
		var LanguageList []string
		for i := 0; i < len(result.Docs[q].BookLanguage); i++{
			LanguageList = append(LanguageList, result.Docs[q].BookLanguage[i])
		}
		for i := 0; i < len(LanguageList); i++{
			fmt.Println("       ",LanguageList[i])
		}

	} else if type_search == "3"{

			printText("\n1) Create simple cli application for finding all works from authors of specific book.\n \n2) Application has to find all authors for book and it will print list of all their works.\n \n3) Create list of works for each author (name, revision).\n \n4) Print result to stdout in yaml format sorted by author name, count of revision (asc, desc as argument). Names of authors have to be part of output.\n \nWrite the title of the book.\n \nExample: Harry Potter\n")
			fmt.Print("Enter a title: ")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			name = strings.ReplaceAll(name, "\n", "")
			origName := name
			if len(origName) > 40{
				fmt.Println("too match word")
				os.Exit(0)
			}
			name = strings.ReplaceAll(name, " ", "%20")
			name = strings.ToLower(name)
			result := GetDataBook("http://openlibrary.org/search.json?title=" + name)
			q := 0
			for i := 0; i < 98; i++{
				if result.Docs[i].Title == origName{
					break
				} else {
					q = q + 1
				}
			}
			fmt.Println("Book:", result.Docs[q].Title)
			fmt.Println("First publish year:", result.Docs[q].FirstPublishYear) 
			fmt.Println("Number of pages:", result.Docs[q].NumberOfPagesMedian)
		
			fmt.Println("Authors: ")
			var AuthorsList []string
			for i := 0; i < len(result.Docs[q].Authors); i++{
				AuthorsList = append(AuthorsList, result.Docs[q].Authors[i])
			}
			sort.Strings(AuthorsList)
			for i := 0; i < len(AuthorsList); i++{
				fmt.Println("       ",AuthorsList[i])
			}
			fmt.Println("Publishers: ")
			var PublishersList []string
			for i := 0; i < len(result.Docs[q].BookPublisher); i++{
				PublishersList = append(PublishersList, result.Docs[q].BookPublisher[i])
			}
			for i := 0; i < len(PublishersList); i++{
				fmt.Println("       ",PublishersList[i])
			}
			fmt.Println("Publish place: ")
			var PublishPlaceList []string
			for i := 0; i < len(result.Docs[q].PublishPlace); i++{
				PublishPlaceList = append(PublishPlaceList, result.Docs[q].PublishPlace[i])
			}
			for i := 0; i < len(PublishPlaceList); i++{
				fmt.Println("       ",PublishPlaceList[i])
			}
			fmt.Println("Language: ")
			var LanguageList []string
			for i := 0; i < len(result.Docs[q].BookLanguage); i++{
				LanguageList = append(LanguageList, result.Docs[q].BookLanguage[i])
			}
			for i := 0; i < len(LanguageList); i++{
				fmt.Println("       ",LanguageList[i])
			}
			fmt.Println("+----------------------------------------------------+\n")
			for i := 0; i < len(AuthorsList); i++{
				fmt.Println("Name:", AuthorsList[i])

				name = strings.ReplaceAll(AuthorsList[i], " ", "%20")
				name = strings.ReplaceAll(name, "\n", "")
				name = strings.ToLower(name)
				resultBook1 := GetDataBooksAuthor(("http://openlibrary.org/search.json?author=" + name))
				var books []string
				var firstPublishYr []int
				for _, rec := range resultBook1.Docs {
						books = append(books, rec.Title)
						firstPublishYr = append(firstPublishYr, rec.FirstPublishYear)
				}
				
				fmt.Println("\nBooks:")
				for i := 0; i < len(books); i++{
					fmt.Println("   ", books[i], firstPublishYr[i])
				}
				fmt.Println("+----------------------------------------------------+\n")
		
			}
			
				
	}

		
	

    


}

func printText(word string){
	splitString := strings.Split(word, "\n")
	var longestString string
	var longestLength int

	for _, s := range splitString {
	  if len(s) > longestLength {
		longestString = s
		longestLength = len(s)
	  }
	}
  
	fmt.Print("+")
	for i := 0; i < len(longestString)+2; i++ {
	  fmt.Print("-")
	}
	fmt.Println("+")

	
	if len(splitString) > 1{
	  for i := 0; i < len(splitString); i++ {
		  fmt.Print("| ")
		  fmt.Print(splitString[i])
		  for q := 0; q < longestLength - len(splitString[i]); q++ {
			  fmt.Print(" ")
		  }
		  fmt.Print(" |")
		  fmt.Print("\n")
		  
	  }
	}

	fmt.Print("+")
	for i := 0; i < len(longestString)+2; i++ {
	  fmt.Print("-")
	}
	fmt.Println("+")
	fmt.Println("")

}
