package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
	"testing"
)

type iW struct {
	i int
	f func() error
}

func (w *iW) Execute() error {
	log.Println(w.i)
	return w.f()
}
func TestNewWorkerPool(t *testing.T) {
	p := NewWorkerPool(10)
	go func() {
		for i := 0; i < 100; i++ {
			w := &iW{i: i, f: func() error { return nil }}
			p.Add(w)
		}
		p.SendDone()
	}()
	p.RunWithBlock()
}

var data = `[
    {
        "id": 18358980,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u7f51\u7edc\u7ea7\u5b89\u5168\u6240\u9762\u4e34\u7684\u4e3b\u8981\u653b\u51fb\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752545",
                "content": "\u7a83\u542c\u3001\u6b3a\u9a97\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752546",
                "content": "\u81ea\u7136\u707e\u5bb3\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752547",
                "content": "\u76d7\u7a83\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752548",
                "content": "\u7f51\u7edc\u5e94\u7528\u8f6f\u4ef6\u7684\u7f3a\u9677\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752548",
        "answer": "65752545",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 10.7,
        "type": 0,
        "isanswerright": false,
        "ordernum": 1,
        "programitems": {}
    },
    {
        "id": 18359000,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5173\u4e8e\u8fde\u63a5\u6570\u636e\u5e93\u7cfb\u7edf\u7684\u5e10\u53f7\uff0c\u4ee5\u4e0b\u8bf4\u6cd5\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; )",
        "answeritems": [
            {
                "sn": "65752587",
                "content": "\u53ef\u4ee5\u4f7f\u7528sa\u7b49\u7ba1\u7406\u5e10\u53f7\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752588",
                "content": "\u53ef\u4ee5\u4f7f\u7528\u9ad8\u7ea7\u522b\u6743\u9650\u5e10\u53f7\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752589",
                "content": "\u5c3d\u53ef\u80fd\u4f7f\u7528\u4f4e\u7ea7\u522b\u6743\u9650\u5e10\u53f7\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752590",
                "content": "\u65e0\u6240\u8c13\uff0c\u9ad8\u7ea7\u522b\u548c\u4f4e\u7ea7\u522b\u5e10\u53f7\u90fd\u884c\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752589",
        "answer": "65752589",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 7.8,
        "type": 0,
        "isanswerright": true,
        "ordernum": 2,
        "programitems": {}
    },
    {
        "id": 18359001,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u4ee5\u4e0b\u54ea\u4e2a\u662f\u7aef\u53e3\u626b\u63cf\u5de5\u5177( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752591",
                "content": "Nmap",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752592",
                "content": "Snoop",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752593",
                "content": "Firewall-2",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752594",
                "content": "NETXRay",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752591",
        "answer": "65752591",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 5.8,
        "type": 0,
        "isanswerright": true,
        "ordernum": 3,
        "programitems": {}
    },
    {
        "id": 18359002,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u4e0b\u5217\u5c5e\u4e8e\u5b89\u5168\u603b\u4f53\u8981\u6c42\u4e2dA\u7c7b\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752595",
                "content": "\u53e3\u4ee4\u5b89\u5168\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752596",
                "content": "\u8bbf\u95ee\u901a\u9053\u63a7\u5236\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752597",
                "content": "web\u5e94\u7528\u5b89\u5168\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752598",
                "content": "\u5b89\u5168\u8d44\u6599\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752598",
        "answer": "65752596",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 114.7,
        "type": 0,
        "isanswerright": false,
        "ordernum": 4,
        "programitems": {}
    },
    {
        "id": 18359003,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u6570\u636e\u4fe1\u606f\u662f\u5426\u88ab\u7be1\u6539\u7531\u54ea\u4e9b\u6280\u672f\u6765\u5224\u65ad( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752599",
                "content": "\u6570\u636e\u5b8c\u6574\u6027\u63a7\u5236\u6280\u672f",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752600",
                "content": "\u8eab\u4efd\u8bc6\u522b\u6280\u672f",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752601",
                "content": "\u8bbf\u95ee\u63a7\u5236\u6280\u672f",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752602",
                "content": "\u5165\u4fb5\u68c0\u6d4b\u6280\u672f",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752599",
        "answer": "65752599",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 12.8,
        "type": 0,
        "isanswerright": true,
        "ordernum": 5,
        "programitems": {}
    },
    {
        "id": 18359004,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5728\u5bf9\u79f0\u5bc6\u94a5\u5bc6\u7801\u4f53\u5236\u4e2d\uff0c\u52a0\u3001\u89e3\u5bc6\u53cc\u65b9\u7684\u5bc6\u94a5( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752603",
                "content": "\u53cc\u65b9\u5404\u81ea\u62e5\u6709\u4e0d\u540c\u7684\u5bc6\u94a5",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752604",
                "content": "\u53cc\u65b9\u7684\u5bc6\u94a5\u53ef\u76f8\u540c\u4e5f\u53ef\u4e0d\u540c",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752605",
                "content": "\u53cc\u65b9\u62e5\u6709\u76f8\u540c\u7684\u5bc6\u94a5",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752606",
                "content": "\u53cc\u65b9\u7684\u5bc6\u94a5\u53ef\u968f\u610f\u6539\u53d8",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752605",
        "answer": "65752605",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 14.5,
        "type": 0,
        "isanswerright": true,
        "ordernum": 6,
        "programitems": {}
    },
    {
        "id": 18359005,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5bf9\u5408\u4f5c\u65b9\u4ea7\u54c1\u6ee1\u8db3\u7f51\u7edc\u5b89\u5168\u603b\u4f53\u8981\u6c42\uff0c\u4e0b\u9762\u90a3\u4e2a\u8bf4\u6cd5\u6b63\u786e\u7684( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752607",
                "content": "\u6240\u6709\u5408\u4f5c\u65b9\u4ea7\u54c1\u73b0\u6709\u7248\u672c\u53ef\u4ee5\u4fdd\u6301\u73b0\u72b6\uff0c\u4e0d\u5bf9\u5b89\u5168\u8981\u6c42\u8fdb\u884c\u6574\u6539\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752608",
                "content": "\u6240\u6709\u5408\u4f5c\u65b9\u4ea7\u54c1\u73b0\u6709\u7248\u672c\u9700\u57282012\u5e74\u4e4b\u524d\u6ee1\u8db3A\u7c7b\u5b89\u5168\u8981\u6c42\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752609",
                "content": "\u6240\u6709\u5408\u4f5c\u65b9\u4ea7\u54c1\u73b0\u6709\u7248\u672c\u9700\u57282013\u5e74\u5e95\u524d\u6ee1\u8db3\u5168\u90e8\u5b89\u5168\u8981\u6c42\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752610",
                "content": "\u6240\u6709\u5408\u4f5c\u65b9\u65b0\u7acb\u9879\u4ea7\u54c1\u7248\u672c\u4f18\u5148\u6ee1\u8db3A\u7c7b\u8981\u6c42\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752610",
        "answer": "65752608",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 75.3,
        "type": 0,
        "isanswerright": false,
        "ordernum": 7,
        "programitems": {}
    },
    {
        "id": 18359006,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5bf9\u7ba1\u7406\u8bbf\u95ee\u901a\u9053\u5b89\u5168\u8981\u6c42\uff0c\u4e0b\u9762\u8bf4\u6cd5\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752611",
                "content": "\u7cfb\u7edf\u7684\u7ba1\u7406\u529f\u80fd\u548c\u4e1a\u52a1\u529f\u80fd\u53ef\u4ee5\u90e8\u7f72\u5728\u540c\u4e00\u53f0\u4e3b\u673a\u4e0a\u7ed1\u5b9a\u540c\u4e00\u4e2a\u7aef\u53e3\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752612",
                "content": "\u53ea\u80fd\u901a\u8fc7\u5b89\u5168\u57df\u5212\u5206\u6765\u63a7\u5236\u6700\u7ec8\u7528\u6237\u4e0d\u80fd\u8bbf\u95ee\u7ba1\u7406\u63a5\u53e3\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752613",
                "content": "\u8bbe\u5907\u5916\u90e8\u53ef\u89c1\u7684\u7ba1\u7406\u8bbf\u95ee\u901a\u9053\u9700\u8981\u6709\u63a5\u5165\u8ba4\u8bc1\u673a\u5236\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752614",
                "content": "\u7cfb\u7edf\u7684\u7ba1\u7406\u529f\u80fd\u548c\u4e1a\u52a1\u529f\u80fd\u5fc5\u987b\u90e8\u7f72\u5728\u4e0d\u540c\u4e3b\u673a\u4e0a\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752613",
        "answer": "65752613",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 33.2,
        "type": 0,
        "isanswerright": true,
        "ordernum": 8,
        "programitems": {}
    },
    {
        "id": 18359007,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5bf9\u7528\u6237\u9762\u4e0e\u7ba1\u7406\u9762\u7684\u9694\u79bb\uff0c\u4e0b\u5217\u54ea\u4e2a\u8bf4\u6cd5\u9519\u8bef\u7684( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752615",
                "content": "\u53ef\u4ee5\u901a\u8fc7ACL\u65b9\u5f0f\u6765\u5b9e\u73b0\u7528\u6237\u9762\u4e0e\u7ba1\u7406\u9762\u7684\u9694\u79bb\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752616",
                "content": "\u53ef\u4ee5\u901a\u8fc7VLAN\u65b9\u5f0f\u6765\u5b9e\u73b0\u7528\u6237\u9762\u4e0e\u7ba1\u7406\u9762\u7684\u9694\u79bb\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752617",
                "content": "\u53ef\u4ee5\u901a\u8fc7\u9632\u706b\u5899\u6765\u5b9e\u73b0\u7528\u6237\u9762\u4e0e\u7ba1\u7406\u9762\u7684\u9694\u79bb\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752618",
                "content": "\u53ef\u4ee5\u901a\u8fc7\u8ba4\u8bc1\u53ca\u6743\u9650\u63a7\u5236\u6765\u5b9e\u73b0\u7528\u6237\u9762\u4e0e\u7ba1\u7406\u9762\u7684\u9694\u79bb\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752617",
        "answer": "65752618",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 64.1,
        "type": 0,
        "isanswerright": false,
        "ordernum": 9,
        "programitems": {}
    },
    {
        "id": 18359008,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5bf9\u64cd\u4f5c\u7cfb\u7edf\u5b89\u5168\uff0c\u4e0b\u9762\u8bf4\u6cd5\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752619",
                "content": "\u64cd\u4f5c\u7cfb\u7edf\u5b89\u5168\u4e3b\u8981\u662f\u9488\u5bf9\u4e3b\u6d41\u901a\u7528\u64cd\u4f5c\u7cfb\u7edf(\u5982Windows\u3001Linux\u7b49)\uff0c\u5bf9\u5d4c\u5165\u5f0f\u64cd\u4f5c\u7cfb\u7edf\u6ca1\u6709\u8981\u6c42\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752620",
                "content": "\u64cd\u4f5c\u7cfb\u7edf\u5b89\u5168\u4e2d\u8981\u6c42\u4f7f\u7528Nessus\u7b49\u6f0f\u6d1e\u626b\u63cf\u8f6f\u4ef6\u8fdb\u884c\u5b89\u5168\u626b\u63cf\uff0c\u4e0d\u5b58\u5728\u9ad8\u98ce\u9669\u7ea7\u522b\u7684\u6f0f\u6d1e\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752621",
                "content": "\u64cd\u4f5c\u7cfb\u7edf\u8fdb\u884c\u4e86\u9632\u75c5\u6bd2\u517c\u5bb9\u6027\u6d4b\u8bd5\u540e\uff0c\u53ef\u4ee5\u4e0d\u8fdb\u884c\u7cfb\u7edf\u52a0\u56fa\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752622",
                "content": "\u5bf9\u4e8e\u5408\u4f5c\u65b9\uff0c\u534e\u4e3a\u516c\u53f8\u63a5\u53d7\u5176\u4f7f\u7528\u4efb\u4f55\u4e00\u6b3e\u4e3b\u6d41\u9632\u75c5\u6bd2\u8f6f\u4ef6\u5bf9\u64cd\u4f5c\u7cfb\u7edf\u7684\u517c\u5bb9\u6027\u6d4b\u8bd5\u7ed3\u679c\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752620",
        "answer": "65752620",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 58.5,
        "type": 0,
        "isanswerright": true,
        "ordernum": 10,
        "programitems": {}
    },
    {
        "id": 18359009,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u4e0b\u5217\u5bf9\u901a\u8baf\u77e9\u9635\u63cf\u8ff0\u9519\u8bef\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752623",
                "content": "\u901a\u4fe1\u77e9\u9635\u4e2d\u6240\u63cf\u8ff0\u7684\u6240\u6709\u7aef\u53e3\u90fd\u662f\u7cfb\u7edf\u8fd0\u884c\u548c\u7ef4\u62a4\u6240\u5fc5\u9700\u7684\uff0c\u4e14\u63cf\u8ff0\u6b63\u786e\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752624",
                "content": "\u901a\u4fe1\u77e9\u9635\u4e2d\u63cf\u8ff0\u7684\u4fa6\u542c\u63a5\u53e3\u987b\u660e\u786e\u9650\u5b9a\u5728\u4e00\u4e2a\u5408\u7406\u7684\u8303\u56f4\u4e4b\u5185\uff0c\u5e76\u4e0e\u5b9e\u9645\u7684\u52a8\u6001\u4fa6\u542c\u63a5\u53e3\u8303\u56f4\u4fdd\u6301\u4e00\u81f4\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752625",
                "content": "\u901a\u8baf\u77e9\u9635\u6587\u6863\u5e76\u5df2\u7eb3\u5165\u5230\u4ea7\u54c1\u8d44\u6599\u6e05\u5355\u4e2d\uff0c\u9700\u8981\u968f\u4ea7\u54c1\u6b63\u5f0f\u53d1\u5e03\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752626",
                "content": "\u901a\u8baf\u77e9\u9635\u6587\u6863\u53ea\u7528\u4e8e\u63d0\u4f9b\u7ed9\u5ba2\u6237\u7528\u4e8e\u7aef\u53e3\u4e1a\u52a1\u7528\u9014\u7684\u81ea\u6211\u6f84\u6e05\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752623",
        "answer": "65752626",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 138.2,
        "type": 0,
        "isanswerright": false,
        "ordernum": 11,
        "programitems": {}
    },
    {
        "id": 18359010,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u4ee5\u4e0b\u5bf9Web\u5b89\u5168\u63cf\u8ff0\u9519\u8bef\u7684\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752627",
                "content": "\u5728Web\u5e94\u7528\u8ba4\u8bc1\u4e2d\uff0c\u53ef\u4ee5\u901a\u8fc7\u9a8c\u8bc1\u7801\u6216\u8005\u591a\u6b21\u8fde\u7eed\u5c1d\u8bd5\u767b\u5f55\u5931\u8d25\u540e\u9501\u5b9a\u5e10\u53f7\u6216IP\u6765\u9632\u66b4\u529b\u7834\u89e3\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752628",
                "content": "\u5982\u91c7\u7528\u591a\u6b21\u8fde\u7eed\u5c1d\u8bd5\u767b\u5f55\u5931\u8d25\u540e\u9501\u5b9a\u5e10\u53f7\u6216IP\u7684\u65b9\u5f0f\uff0c\u9700\u652f\u6301\u8fde\u7eed\u767b\u5f55\u5931\u8d25\u9501\u5b9a\u7b56\u7565\u7684\u201c\u5141\u8bb8\u8fde\u7eed\u5931\u8d25\u7684\u6b21\u6570\u201d\u53ef\u56fa\u5b9a\u4e3a\u4e00\u4e2a\u5408\u7406\u7684\u6b21\u6570\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752629",
                "content": "\u6570\u636e\u5728\u8f93\u51fa\u5230\u5ba2\u6237\u7aef\u524d\u5fc5\u987b\u5148\u8fdb\u884cHTML\u7f16\u7801\uff0c\u4ee5\u9632\u6b62\u6267\u884c\u6076\u610f\u4ee3\u7801\u3001\u8de8\u7ad9\u811a\u672c\u653b\u51fb\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752630",
                "content": "\u7528\u6237\u4ea7\u751f\u7684\u6570\u636e\u8981\u5728\u670d\u52a1\u7aef\u8fdb\u884c\u6821\u9a8c\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752629",
        "answer": "65752628",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 56.3,
        "type": 0,
        "isanswerright": false,
        "ordernum": 12,
        "programitems": {}
    },
    {
        "id": 18359011,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u4ee5\u4e0b\u5173\u4e8e\u6709\u52a0\u5bc6\u7b97\u6cd5\u53ca\u5bc6\u94a5\u63cf\u8ff0\uff0c\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752631",
                "content": "\u4ea7\u54c1\u53ef\u4ee5\u6839\u636e\u9700\u8981\u91c7\u7528\u79c1\u6709\u52a0\u5bc6\u7b97\u6cd5\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752632",
                "content": "AES\u662f\u5c5e\u4e8e\u975e\u5bf9\u79f0\u52a0\u5bc6\u7b97\u6cd5\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752633",
                "content": "\u5728\u654f\u611f\u6570\u636e\u7684\u5b89\u5168\u4f20\u8f93\u4e0a\uff0c\u4f18\u5148\u4f7f\u7528\u4e1a\u754c\u7684\u6807\u51c6\u5b89\u5168\u534f\u8bae\uff0c\u5e76\u786e\u4fdd\u5bc6\u94a5\u53ef\u914d\u7f6e\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752634",
                "content": "\u52a0\u5bc6\u7b97\u6cd5\u4e2d\u4f7f\u7528\u7684\u5bc6\u94a5\uff0c\u53ef\u4ee5\u786c\u7f16\u7801\u5728\u4ee3\u7801\u4e2d\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752631",
        "answer": "65752633",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 21.3,
        "type": 0,
        "isanswerright": false,
        "ordernum": 13,
        "programitems": {}
    },
    {
        "id": 18359012,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u672a\u516c\u5f00\u63a5\u53e3\u5f15\u8d77\u7684\u98ce\u9669\uff0c\u4ee5\u4e0b\u8bf4\u6cd5\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752635",
                "content": "\u672a\u516c\u5f00\u63a5\u53e3\u5bb9\u6613\u88ab\u5ba2\u6237\u8d28\u7591\u4e3a\u4ea7\u54c1\u540e\u95e8\uff0c\u4e14\u5bb9\u6613\u88ab\u6076\u610f\u653b\u51fb\u8005\u5229\u7528\uff0c\u4ece\u800c\u589e\u52a0\u6211\u53f8\u4ea7\u54c1\u7684\u5b89\u5168\u98ce\u9669\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752636",
                "content": "\u672a\u516c\u5f00\u63a5\u53e3\u53ea\u8981\u4e0d\u5728\u8d44\u6599\u4e2d\u8bf4\u660e\uff0c\u5ba2\u6237\u4e5f\u4e0d\u77e5\u9053\u600e\u4e48\u7528\uff0c\u4e0d\u4f1a\u7ed9\u4ea7\u54c1\u5e26\u6765\u4efb\u4f55\u98ce\u9669\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752637",
                "content": "\u53ea\u8981\u52a0\u5f3a\u672a\u516c\u5f00\u63a5\u53e3\u7684\u8bbf\u95ee\u63a7\u5236\uff0c\u5373\u4f7f\u672a\u6587\u6863\u5316\u4ea7\u54c1\u4e5f\u4e0d\u4f1a\u6709\u4efb\u4f55\u5f71\u54cd\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752638",
                "content": "\u53ea\u8981\u4ea7\u54c1\u4e0d\u5728\u654f\u611f\u5e02\u573a\u4f7f\u7528\uff0c\u672a\u516c\u5f00\u63a5\u53e3\u4e0d\u4f1a\u7ed9\u4ea7\u54c1\u6216\u516c\u53f8\u5e26\u6765\u98ce\u9669\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752635",
        "answer": "65752635",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 18.9,
        "type": 0,
        "isanswerright": true,
        "ordernum": 14,
        "programitems": {}
    },
    {
        "id": 18359013,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u654f\u611f\u6570\u636e\u4fdd\u62a4\u4e3b\u8981\u662f\u52a0\u5f3a\u54ea\u4e9b\u65b9\u9762\u7684\u4fdd\u62a4( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752639",
                "content": "\u901a\u8fc7\u52a0\u5bc6\u4fdd\u62a4\u52a0\u5f3a\u7cfb\u7edf\u5bf9\u654f\u611f\u6570\u636e\u7684\u5b58\u50a8\u7684\u5b89\u5168\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752640",
                "content": "\u901a\u8fc7\u8ba4\u8bc1\u3001\u6388\u6743\u548c\u52a0\u5bc6\u673a\u5236\u52a0\u5f3a\u654f\u611f\u6570\u636e\u7684\u8bbf\u95ee\u5b89\u5168\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752641",
                "content": "\u901a\u8fc7\u5b89\u5168\u4f20\u8f93\u901a\u9053\u52a0\u5f3a\u654f\u611f\u6570\u636e\u5728\u975e\u4fe1\u4efb\u7f51\u7edc\u4e4b\u95f4\u8fdb\u884c\u4f20\u8f93\u5b89\u5168\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752642",
                "content": "\u4ee5\u4e0a\u5168\u662f",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752642",
        "answer": "65752642",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 18.5,
        "type": 0,
        "isanswerright": true,
        "ordernum": 15,
        "programitems": {}
    },
    {
        "id": 18359014,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5bf9\u7cfb\u7edf\u7ba1\u7406\u53ca\u7ef4\u62a4\u5b89\u5168\u63cf\u8ff0\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752643",
                "content": "\u7cfb\u7edf\u81ea\u8eab\u64cd\u4f5c\u7ef4\u62a4\u7c7b\u53e3\u4ee4\u6ee1\u8db3\u201c\u53e3\u4ee4\u5b89\u5168\u8981\u6c42\u201d\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752644",
                "content": "\u7ba1\u7406\u9762\u6240\u6709\u7684\u7528\u6237\u6d3b\u52a8\u3001\u64cd\u4f5c\u6307\u4ee4\u5fc5\u987b\u8bb0\u5f55\u65e5\u5fd7\uff0c\u65e5\u5fd7\u5185\u5bb9\u8981\u80fd\u652f\u6491\u4e8b\u540e\u7684\u5ba1\u8ba1\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752645",
                "content": "\u7cfb\u7edf\u7684\u7ba1\u7406\u5e73\u9762\u548c\u8fd1\u7aef\u7ef4\u62a4\u7ec8\u7aef\u3001\u7f51\u7ba1\u7ef4\u62a4\u7ec8\u7aef\u95f4\uff0c\u652f\u6301\u4f7f\u7528\u5408\u9002\u7684\u5b89\u5168\u534f\u8bae\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752646",
                "content": "\u4ee5\u4e0a\u5168\u662f",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752646",
        "answer": "65752646",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 24.2,
        "type": 0,
        "isanswerright": true,
        "ordernum": 16,
        "programitems": {}
    },
    {
        "id": 18359015,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5173\u4e8e\u53e3\u4ee4\u590d\u6742\u5ea6\u8981\u6c42\uff0c\u4ee5\u4e0b\u8bf4\u6cd5\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752647",
                "content": "\u666e\u901a\u7528\u6237\u53ca\u7279\u6743\u7528\u6237\u7684\u53e3\u4ee4\u957f\u5ea6\u8981\u6c42\u90fd\u662f\u4e00\u6837\u7684\uff0c\u81f3\u5c116\u4e2a\u5b57\u7b26\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752648",
                "content": "\u53e3\u4ee4\u4e0d\u80fd\u548c\u5e10\u53f7\u6216\u8005\u5e10\u53f7\u7684\u9006\u5e8f\u76f8\u540c\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752649",
                "content": "\u53e3\u4ee4\u4e0d\u80fd\u5305\u542b\u7a7a\u683c\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752650",
                "content": "\u53e3\u4ee4\u4e0d\u80fd\u5305\u542b\u6570\u5b57\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752648",
        "answer": "65752648",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 17.1,
        "type": 0,
        "isanswerright": true,
        "ordernum": 17,
        "programitems": {}
    },
    {
        "id": 18359016,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5bf9\u4e8e\u65e5\u5fd7\u5b89\u5168\u8bbe\u8ba1\u89c4\u5219\u63cf\u8ff0\u4e2d\uff0c\u4e0d\u5305\u62ec\u4e0b\u5217\u90a3\u4e00\u9879 ( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752651",
                "content": "\u7cfb\u7edf\u5fc5\u987b\u5bf9\u5b89\u5168\u4e8b\u4ef6\u53ca\u64cd\u4f5c\u4e8b\u4ef6\u8fdb\u884c\u65e5\u5fd7\u8bb0\u5f55",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752652",
                "content": "\u5e94\u8be5\u5f00\u653e\u5bf9\u5b89\u5168\u65e5\u5fd7\u7684\u8bbf\u95ee",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752653",
                "content": "\u5bf9\u65e5\u5fd7\u6a21\u5757\u5360\u6709\u8d44\u6e90\u6709\u53ef\u914d\u7f6e\u7684\u9650\u5236\u673a\u5236",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752654",
                "content": "\u5b89\u5168\u4e8b\u4ef6\u7684\u7ed3\u679c\uff0c\u4e0d\u7ba1\u6210\u529f\u8fd8\u662f\u5931\u8d25\uff0c\u90fd\u8981\u8bb0\u5f55\u65e5\u5fd7",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752652",
        "answer": "65752652",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 25.5,
        "type": 0,
        "isanswerright": true,
        "ordernum": 18,
        "programitems": {}
    },
    {
        "id": 18359017,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u5bf9\u65e5\u5fd7\u5ba1\u8ba1\u63cf\u8ff0\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752655",
                "content": "\u65e5\u5fd7\u8bb0\u5f55\u662f\u5b89\u5168\u4e8b\u4ef6\u4e2d\u4e8b\u540e\u8ffd\u6eaf\uff0c\u5b9a\u4f4d\u95ee\u9898\u539f\u56e0\u53ca\u5212\u5206\u4e8b\u6545\u8d23\u4efb\u7684\u91cd\u8981\u624b\u6bb5\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752656",
                "content": "\u8be6\u7ec6\u7684\u65e5\u5fd7\u8bb0\u5f55\u4e00\u65b9\u9762\u53ef\u4ee5\u5e2e\u52a9\u534e\u4e3a\u6487\u6e05\u4e0e\u6211\u65b9\u4e0d\u76f8\u5173\u7684\u8d23\u4efb\uff08\u5982\u8fd0\u8425\u5546\u7684\u5931\u8bef\uff09\uff0c\u4e5f\u53ef\u4ee5\u5e2e\u52a9\u56de\u6eaf\u5386\u53f2\u64cd\u4f5c\uff0c\u63d0\u9ad8\u73b0\u7f51\u95ee\u9898\u5b9a\u4f4d\u7684\u6548\u7387\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752657",
                "content": "\u7ba1\u7406\u9762\u6240\u6709\u7684\u7528\u6237\u6d3b\u52a8\u3001\u64cd\u4f5c\u6307\u4ee4\u5fc5\u987b\u8bb0\u5f55\u65e5\u5fd7\uff0c\u65e5\u5fd7\u5185\u5bb9\u8981\u80fd\u652f\u6491\u4e8b\u540e\u7684\u5ba1\u8ba1\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752658",
                "content": "\u4ee5\u4e0a\u90fd\u662f\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752658",
        "answer": "65752658",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 12.0,
        "type": 0,
        "isanswerright": true,
        "ordernum": 19,
        "programitems": {}
    },
    {
        "id": 18359018,
        "user": "icollege-migrate",
        "way": 0,
        "question": "\u4ee5\u4e0b\u5bf9\u654f\u611f\u6570\u636e\u5b58\u50a8\u91c7\u53d6\u7684\u63aa\u65bd\uff0c\u4e0d\u5408\u7406\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752659",
                "content": "\u7981\u6b62\u5c06\u654f\u611f\u6570\u636e\u5b58\u50a8\u5728\u4ee3\u7801\u4e2d\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752660",
                "content": "\u5c06\u5bc6\u94a5\u6216\u53e3\u4ee4\u5b58\u653e\u5728\u6570\u636e\u5e93\u6587\u4ef6\u4e2d\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752661",
                "content": "\u7981\u6b62\u5728cookie\u4e2d\u4ee5\u660e\u6587\u5b58\u653e\u654f\u611f\u6570\u636e\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752662",
                "content": "\u7981\u6b62\u5c06\u654f\u611f\u6570\u636e\u4ee5\u660e\u6587\u5b58\u653e\u5728\u9690\u85cf\u57df\u4e2d\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752660",
        "answer": "65752660",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 24.1,
        "type": 0,
        "isanswerright": true,
        "ordernum": 20,
        "programitems": {}
    },
    {
        "id": 18359019,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u672a\u516c\u5f00\u63a5\u53e3\u4e3b\u8981\u6307\u4ee5\u4e0b\u54ea\u51e0\u7c7b\u63a5\u53e3( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752663",
                "content": "\u7ed5\u8fc7\u5b89\u5168\u9274\u6743\u673a\u5236\u8bbf\u95ee\u7cfb\u7edf\u7684\u63a5\u53e3\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752664",
                "content": "\u56e0\u53d7\u9650\u4f7f\u7528\u800c\u9690\u85cf\u6216\u672a\u6587\u6863\u5316\u7684\u547d\u4ee4/\u53c2\u6570/\u7aef\u53e3\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752665",
                "content": "\u7528\u4e8e\u7cfb\u7edf\u7ba1\u7406\u529f\u80fd\u7684\u63a5\u53e3\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752666",
                "content": "\u7528\u4e8e\u4ea7\u54c1\u5185\u90e8\u6d4b\u8bd5\u9636\u6bb5\u4f7f\u7528\uff0c\u4f46\u5728\u6b63\u5f0f\u53d1\u5e03\u4e2d\u5df2\u5220\u9664\u7684\u63a5\u53e3\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752663,65752664,65752666",
        "answer": "65752663,65752664",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 66.4,
        "type": 1,
        "isanswerright": false,
        "ordernum": 21,
        "programitems": {}
    },
    {
        "id": 18359020,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4ee5\u4e0b\u9488\u5bf9\u6570\u636e\u5e93\u5b89\u5168\u8bf4\u6cd5\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752667",
                "content": "\u6570\u636e\u5e93\u82e5\u5b58\u5728\u591a\u4e2a\u9ed8\u8ba4\u5e10\u53f7\uff0c\u5fc5\u987b\u5c06\u4e0d\u4f7f\u7528\u7684\u5e10\u53f7\u7981\u7528\u6216\u5220\u9664\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752668",
                "content": "\u6570\u636e\u5e93\u53e3\u4ee4\u53ef\u4ee5\u4f7f\u7528\u6570\u636e\u5e93\u5382\u5546\u7684\u7f3a\u7701\u53e3\u4ee4\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752669",
                "content": "\u4f7f\u7528\u5355\u72ec\u7684\u64cd\u4f5c\u7cfb\u7edf\u5e10\u53f7\u6765\u8fd0\u884c\u6570\u636e\u5e93\uff0c\u6570\u636e\u5e93\u4e2d\u7684\u654f\u611f\u6587\u4ef6\uff0c\u9700\u8981\u4e25\u683c\u63a7\u5236\u8bbf\u95ee\u6743\u9650\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752670",
                "content": "\u4f7f\u7528\u4e3b\u6d41\u7684\u7cfb\u7edf\u626b\u63cf\u8f6f\u4ef6\u8fdb\u884c\u5b89\u5168\u626b\u63cf\uff0c\u4e0d\u5b58\u5728\u201c\u9ad8\u201d\u7ea7\u522b\u7684\u6f0f\u6d1e\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752667,65752669",
        "answer": "65752667,65752669,65752670",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 129.9,
        "type": 1,
        "isanswerright": false,
        "ordernum": 22,
        "programitems": {}
    },
    {
        "id": 18359021,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u5e94\u7528\u5c42\u4e3b\u8981\u7684\u5b89\u5168\u9700\u6c42\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752671",
                "content": "\u8eab\u4efd\u8ba4\u8bc1",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752672",
                "content": "\u9632\u6b62\u7269\u7406\u7834\u574f",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752673",
                "content": "\u8bbf\u95ee\u63a7\u5236",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752674",
                "content": "\u9632\u6b62\u4eba\u4e3a\u653b\u51fb",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752671,65752673,65752674",
        "answer": "65752671,65752673",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 48.5,
        "type": 1,
        "isanswerright": false,
        "ordernum": 23,
        "programitems": {}
    },
    {
        "id": 18359022,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u7ba1\u7406\u5e73\u9762\u7684\u63a5\u53e3\u7684\u8bbf\u95ee\u4fdd\u62a4\u4e3b\u8981\u6709\u54ea\u4e9b\u8981\u6c42( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752675",
                "content": "\u7cfb\u7edf\u7684\u7ba1\u7406\u529f\u80fd\u548c\u4e1a\u52a1\u529f\u80fd\u5e94\u80fd\u591f\u5206\u522b\u90e8\u7f72\u5728\u4e0d\u540c\u7684\u4e3b\u673a\u4e0a\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752676",
                "content": "\u7cfb\u7edf\u7684\u7ba1\u7406\u529f\u80fd\u548c\u4e1a\u52a1\u529f\u80fd\u5982\u679c\u90e8\u7f72\u5728\u540c\u4e00\u4e3b\u673a\u4e0a\uff0c\u5e94\u80fd\u591f\u5206\u522b\u7ed1\u5b9a\u4e0d\u540c\u7684IP\u5730\u5740\u6216\u4e0d\u540c\u7684\u7aef\u53e3\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752677",
                "content": "\u901a\u8fc7\u5b89\u5168\u57df\u5212\u5206\u3001\u9632\u706b\u5899\u8bbf\u95ee\u63a7\u5236\uff0c\u6700\u7ec8\u7528\u6237\u4e0d\u80fd\u8bbf\u95ee\u7ba1\u7406\u63a5\u53e3\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752678",
                "content": "\u6240\u6709\u80fd\u5bf9\u7cfb\u7edf\u8fdb\u884c\u7ba1\u7406\u7684\u903b\u8f91\u901a\u4fe1\u7aef\u53e3\u53ca\u534f\u8bae\u90fd\u9700\u5177\u5907\u63a5\u5165\u8ba4\u8bc1\u673a\u5236\uff08\u534f\u8bae\u6807\u51c6\u5b9a\u4e49\u4e2d\u65e0\u8ba4\u8bc1\u673a\u5236\u7684\u9664\u5916\uff09\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752675,65752676,65752678",
        "answer": "65752675,65752676,65752677,65752678",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 246.7,
        "type": 1,
        "isanswerright": false,
        "ordernum": 24,
        "programitems": {}
    },
    {
        "id": 18359023,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4ee5\u4e0b\u5173\u4e8e\u901a\u8baf\u77e9\u9635\u63cf\u8ff0\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752679",
                "content": "\u901a\u8baf\u77e9\u9635\u53ef\u4ee5\u6307\u5bfc\u73b0\u7f51\u9632\u706b\u5899\u7684\u914d\u7f6e\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752680",
                "content": "\u901a\u8baf\u77e9\u9635\u8fd8\u53ef\u4ee5\u7528\u4f5c\u4ea7\u54c1\u7aef\u53e3\u4e1a\u52a1\u7528\u9014\u7684\u81ea\u6211\u6f84\u6e05\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752681",
                "content": "\u4e1a\u52a1\u90e8\u7f72\u5f62\u6001\u4e0d\u4e00\u6837\uff0c\u914d\u7f6e\u4e5f\u4e0d\u4e00\u6837\uff0c\u5f00\u653e\u7684\u7aef\u53e3\u4e5f\u4e0d\u4e00\u6837\uff0c\u901a\u8baf\u77e9\u9635\u4ec5\u4ec5\u9700\u8981\u8bb0\u5f55\u4e1a\u52a1\u7cfb\u7edf\u5bf9\u5916\u7684\u7aef\u53e3\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752682",
                "content": "\u52a8\u6001\u76d1\u542c\u63a5\u53e3\u53ef\u4ee5\u4e0d\u5728\u901a\u8baf\u77e9\u9635\u4e2d\u8fdb\u884c\u63cf\u8ff0\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752680,65752681",
        "answer": "65752679,65752680",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 111.2,
        "type": 1,
        "isanswerright": false,
        "ordernum": 25,
        "programitems": {}
    },
    {
        "id": 18359024,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u5bf9\u7cfb\u7edf\u6240\u6709\u5bf9\u5916\u901a\u8baf\u8fde\u63a5\u5fc5\u987b\u662f\u7cfb\u7edf\u5fc5\u987b\u7684\u7406\u89e3\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752683",
                "content": "\u7cfb\u7edf\u6307\u4ea4\u4ed8\u7ed9\u5ba2\u6237\u8fd0\u884c\u7684\u6574\u4f53\u7cfb\u7edf\uff0c\u5305\u62ec\u81ea\u7814\u7684\u8f6f\u4ef6\u3001\u8f6f\u4ef6\u8fd0\u884c\u7684\u64cd\u4f5c\u7cfb\u7edf\u53ca\u5e94\u7528\u670d\u52a1\u5728\u5185\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752684",
                "content": "\u4ea7\u54c1\u5728\u8bbe\u8ba1\u65f6\u8981\u9075\u5faa\u7aef\u53e3\u6700\u5c0f\u5f00\u653e\u539f\u5219\uff0c\u975e\u4e1a\u52a1\u4ee5\u53ca\u7ef4\u62a4\u9700\u8981\u7684\u7aef\u53e3\u5e94\u9ed8\u8ba4\u5173\u95ed\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752685",
                "content": "\u5e73\u53f0\u5e94\u901a\u8fc7\u89e3\u8026\u964d\u4f4e\u7aef\u53e3\u95f4\u7684\u8026\u5408\u5173\u7cfb\uff0c\u4ee5\u4fbf\u4ea7\u54c1\u6309\u9700\u9009\u62e9\u7aef\u53e3\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752686",
                "content": "\u7cfb\u7edf\u5bf9\u5916\u901a\u8baf\u8fde\u63a5\u662f\u6307\u57fa\u4e8eTCP\u7b49\u8fde\u63a5\u534f\u8bae\u5efa\u7acb\u7684\u7ba1\u7406\u8fde\u63a5\u3002",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752684,65752685,65752686",
        "answer": "65752683,65752684,65752685",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 509.7,
        "type": 1,
        "isanswerright": false,
        "ordernum": 26,
        "programitems": {}
    },
    {
        "id": 18359025,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4e0b\u5217\u90a3\u4e9b\u884c\u4e3a\u662f\u5c5e\u4e8e\u5bf9\u64cd\u4f5c\u7cfb\u7edf\u8fdb\u884c\u5b89\u5168\u4fdd\u62a4\u7684\uff1f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752687",
                "content": "\u4f7f\u7528\u4e3b\u6d41\u6f0f\u6d1e\u626b\u63cf\u8f6f\u4ef6\u5bf9\u64cd\u4f5c\u7cfb\u7edf\u8fdb\u884c\u5b89\u5168\u626b\u63cf\uff0c\u4e0d\u5b58\u5728\u9ad8\u98ce\u9669\u7ea7\u522b\u7684\u6f0f\u6d1e\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752688",
                "content": "\u5bf9\u64cd\u4f5c\u7cfb\u7edf\u8fdb\u884c\u52a0\u56fa\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752689",
                "content": "\u5bf9\u64cd\u4f5c\u7cfb\u7edf\u8fdb\u884c\u9632\u75c5\u6bd2\u626b\u63cf\u53ca\u517c\u5bb9\u6027\u6d4b\u8bd5\uff0c\u626b\u63cf\u7ed3\u679c\u968f\u7248\u672c\u53d1\u5e03\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752690",
                "content": "\u64cd\u4f5c\u7cfb\u7edf\u6253\u5b89\u5168\u8865\u4e01\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752688,65752689,65752690",
        "answer": "65752687,65752688,65752689,65752690",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 64.9,
        "type": 1,
        "isanswerright": false,
        "ordernum": 27,
        "programitems": {}
    },
    {
        "id": 18359026,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4e0b\u5217\u54ea\u4e9b\u5c5e\u4e8eWeb\u5b89\u5168\u8981\u6c42\uff1f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752691",
                "content": "\u8ba4\u8bc1\u6a21\u5757\u5fc5\u987b\u91c7\u7528\u9632\u66b4\u529b\u7834\u89e3\u673a\u5236\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752692",
                "content": "\u5bf9\u4e8e\u6bcf\u4e00\u4e2a\u9700\u8981\u6388\u6743\u8bbf\u95ee\u7684\u9875\u9762\u6216servlet\u7684\u8bf7\u6c42\u90fd\u5fc5\u987b\u6838\u5b9e\u7528\u6237\u7684\u4f1a\u8bdd\u6807\u8bc6\u662f\u5426\u5408\u6cd5\u3001\u7528\u6237\u662f\u5426\u88ab\u6388\u6743\u6267\u884c\u8fd9\u4e2a\u64cd\u4f5c\uff0c\u4ee5\u9632\u6b62URL\u8d8a\u6743\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752693",
                "content": "\u767b\u5f55\u8fc7\u7a0b\u4e2d\uff0c\u5f80\u670d\u52a1\u5668\u7aef\u4f20\u9012\u7528\u6237\u540d\u548c\u53e3\u4ee4\u65f6\uff0c\u5fc5\u987b\u91c7\u7528HTTPS\u5b89\u5168\u534f\u8bae\u4e5f\u5c31\u662f\u5e26\u670d\u52a1\u5668\u7aef\u8bc1\u4e66\u7684SSL\uff09\uff0c\u53ea\u63d0\u4f9b\u672c\u673a\u63a5\u5165\u3001\u767b\u5f55\uff0c\u505a\u8bbe\u5907\u7ba1\u7406\u4f7f\u7528\u7684\u573a\u666f\u6682\u65f6\u4e0d\u8981\u6c42\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752694",
                "content": "\u4f7f\u7528\u4e3b\u6d41Web\u5b89\u5168\u626b\u63cf\u5de5\u5177\u626b\u63cfWeb\u670d\u52a1\u5668\u548cWeb\u5e94\u7528\uff0c\u4e0d\u5b58\u5728\u201c\u9ad8\u201d\u7ea7\u522b\u7684\u6f0f\u6d1e\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752691,65752692,65752693",
        "answer": "65752691,65752692,65752693,65752694",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 262.4,
        "type": 1,
        "isanswerright": false,
        "ordernum": 28,
        "programitems": {}
    },
    {
        "id": 18359027,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u5bf9\u7528\u6237\u4ea7\u751f\u7684\u6570\u636e\u5fc5\u987b\u5728\u670d\u52a1\u7aef\u8fdb\u884c\u6821\u9a8c\uff0c\u4e0b\u9762\u8bf4\u6cd5\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752695",
                "content": "\u5982\u679c\u5bf9\u7528\u6237\u4ea7\u751f\u7684\u6570\u636e\u653e\u5230\u5ba2\u6237\u7aef\u6821\u9a8c\uff0c\u5bb9\u6613\u88ab\u7ed5\u8fc7\uff0c\u6821\u9a8c\u5982\u540c\u865a\u8bbe\uff0c\u6076\u610f\u7528\u6237\u53ef\u4ee5\u968f\u610f\u6784\u9020\u6570\u636e\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752696",
                "content": "\u9632\u6b62\u8de8\u7ad9\u811a\u672c\u653b\u51fb\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752697",
                "content": "\u9632\u6b62\u6267\u884c\u6076\u610f\u4ee3\u7801\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752698",
                "content": "\u9632\u6b62SQL\u6ce8\u5165\u3001\u547d\u4ee4\u6ce8\u5165\u3001\u7f13\u51b2\u533a\u6ea2\u51fa\u7b49\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752695,65752696,65752697,65752698",
        "answer": "65752695,65752696,65752697,65752698",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 173.6,
        "type": 1,
        "isanswerright": true,
        "ordernum": 29,
        "programitems": {}
    },
    {
        "id": 18359028,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4e0b\u5217\u534f\u8bae\u54ea\u4e9b\u5c5e\u4e8e\u5b89\u5168\u8bbf\u95ee\u534f\u8bae\uff1f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752699",
                "content": "HTTP",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752700",
                "content": "SFTP",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752701",
                "content": "IPSec",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752702",
                "content": "SSL",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752700,65752702",
        "answer": "65752700,65752701,65752702",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 7.0,
        "type": 1,
        "isanswerright": false,
        "ordernum": 30,
        "programitems": {}
    },
    {
        "id": 18359029,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4ea7\u54c1\u5f00\u53d1\u3001\u53d1\u5e03\u548c\u5b89\u88c5\u5b89\u5168\u8981\u6c42\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752703",
                "content": "\u7981\u6b62\u5b58\u5728\u4efb\u4f55\u201c\u672a\u516c\u5f00\u63a5\u53e3\u201d\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752704",
                "content": "\u5408\u4f5c\u65b9\u9700\u5728\u4ea7\u54c1\u8d44\u6599\u4e2d\u516c\u5f00\u4eba\u673a\u4ea4\u4e92\u63a5\u53e3\u3001\u4e0e\u7b2c\u4e09\u65b9\u7cfb\u7edf\u5bf9\u63a5\u63a5\u53e3\uff08\u63a8\u8350\u901a\u8fc7\u6280\u672f\u624b\u6bb5\u9650\u5236\u7b2c\u4e09\u65b9\u53ea\u80fd\u8bbf\u95ee\u4e1a\u52a1\u5fc5\u987b\u7684\u7aef\u53e3\uff09\u3001\u9700\u7ecf\u5e38\u4eba\u5de5\u6539\u52a8\u7684\u914d\u7f6e\u6587\u4ef6\u7b49\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752705",
                "content": "\u5185\u90e8\u901a\u4fe1\u7684\u673a\u673a\u63a5\u53e3\u5e94\u80fd\u591f\u901a\u8fc7\u5b89\u5168\u7684\u7efc\u5408\u624b\u6bb5\u4fdd\u969c\uff08\u5982\u7ec4\u7f51\u7b49\uff09\uff0c\u4e0d\u9700\u8981\u516c\u5f00\uff0c\u4f46\u9700\u8981\u5728\u8d44\u6599\u8bf4\u8bf4\u660e\u5b89\u5168\u4fdd\u969c\u624b\u6bb5\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752706",
                "content": "\u5728\u8f6f\u4ef6\u5305\uff08\u542b\u8865\u4e01\u5305\uff09\u53d1\u5e03\u524d\uff0c\u9700\u8981\u7ecf\u8fc7\u81f3\u5c11\u4e00\u6b3e\u4e3b\u6d41\u9632\u75c5\u6bd2\u8f6f\u4ef6\u626b\u63cf\uff0c\u4fdd\u8bc1\u9632\u75c5\u6bd2\u8f6f\u4ef6\u4e0d\u4ea7\u751f\u544a\u8b66\uff0c\u7279\u6b8a\u60c5\u51b5\u4e0b\u5bf9\u544a\u8b66\u4f5c\u51fa\u89e3\u91ca\u8bf4\u660e\u3002\u626b\u63cf\u8bb0\u5f55\uff08\u9632\u75c5\u6bd2\u8f6f\u4ef6\u540d\u79f0\u3001\u8f6f\u4ef6\u7248\u672c\u3001\u75c5\u6bd2\u5e93\u7248\u672c\u3001\u626b\u63cf\u65f6\u95f4\u3001\u626b\u63cf\u7ed3\u679c\u7b49\uff09\u5b58\u6863\u5e76\u968f\u8f6f\u4ef6\u5305\uff08\u542b\u8865\u4e01\u5305\uff09\u53d1\u5e03\u7ed9\u5ba2\u6237\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752703,65752704,65752705,65752706",
        "answer": "65752703,65752704,65752705,65752706",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 27.7,
        "type": 1,
        "isanswerright": true,
        "ordernum": 31,
        "programitems": {}
    },
    {
        "id": 18359030,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4e0b\u9762\u54ea\u4e9b\u5b89\u5168\u4fdd\u62a4\u662f\u9488\u5bf9\u654f\u611f\u6570\u636e\u7684\uff1f( &nbsp; &nbsp; &nbsp; &nbsp;)",
        "answeritems": [
            {
                "sn": "65752707",
                "content": "\u53e3\u4ee4\u4e0d\u660e\u6587\u5b58\u50a8\u5728\u7cfb\u7edf\u4e2d\uff0c\u901a\u8fc7\u52a0\u5bc6\u8fdb\u884c\u4fdd\u62a4\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752708",
                "content": "\u5bf9\u94f6\u884c\u8d26\u53f7\u7b49\u654f\u611f\u6570\u636e\u7684\u8bbf\u95ee\u8981\u6709\u8ba4\u8bc1\u3001\u6388\u6743\u548c\u52a0\u5bc6\u673a\u5236\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752709",
                "content": "\u5728\u975e\u4fe1\u4efb\u7f51\u7edc\u4e4b\u95f4\u8fdb\u884c\u654f\u611f\u6570\u636e\uff08\u5305\u62ec\u53e3\u4ee4\uff0c\u94f6\u884c\u5e10\u53f7\uff0c\u6279\u91cf\u4e2a\u4eba\u6570\u636e\u7b49\uff09\u7684\u4f20\u8f93\u987b\u91c7\u7528\u5b89\u5168\u4f20\u8f93\u901a\u9053\u6216\u8005\u52a0\u5bc6\u540e\u4f20\u8f93\uff0c\u6709\u6807\u51c6\u534f\u8bae\u89c4\u5b9a\u9664\u5916\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752710",
                "content": "\u7cfb\u7edf\u7684\u7ba1\u7406\u5e73\u9762\u548c\u8fd1\u7aef\u7ef4\u62a4\u7ec8\u7aef(\u5982LMT)\u3001\u7f51\u7ba1\u7ef4\u62a4\u7ec8\u7aef\u95f4\uff0c\u652f\u6301\u4f7f\u7528\u5408\u9002\u7684\u5b89\u5168\u534f\u8bae\u8fdb\u884c\u901a\u4fe1\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752707,65752708,65752709",
        "answer": "65752707,65752708,65752709,65752710",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 164.0,
        "type": 1,
        "isanswerright": false,
        "ordernum": 32,
        "programitems": {}
    },
    {
        "id": 18359031,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4e0b\u5217\u54ea\u4e9b\u5c5e\u4e8e\u9700\u8981\u8bb0\u5f55\u5728\u65e5\u5fd7\u4e2d\u7684\u7ba1\u7406\u5c42\u6d3b\u52a8\uff1f( &nbsp; &nbsp; &nbsp; &nbsp;)",
        "answeritems": [
            {
                "sn": "65752711",
                "content": "\u767b\u5f55\u548c\u6ce8\u9500",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752712",
                "content": "\u7528\u6237\u7684\u9501\u5b9a\u548c\u89e3\u9501\uff0c\u7981\u7528\u548c\u6062\u590d",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752713",
                "content": "\u5bf9\u7cfb\u7edf\u8fdb\u884c\u542f\u52a8\u3001\u5173\u95ed\u3001\u91cd\u542f\u3001\u6682\u505c\u3001\u6062\u590d\u3001\u5012\u6362",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752714",
                "content": "\u6240\u6709\u5e10\u6237\u7684\u547d\u4ee4\u884c\u64cd\u4f5c\u547d\u4ee4",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752711,65752712,65752713,65752714",
        "answer": "65752711,65752712,65752713,65752714",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 23.5,
        "type": 1,
        "isanswerright": true,
        "ordernum": 33,
        "programitems": {}
    },
    {
        "id": 18359032,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4e0b\u5217\u5c5e\u4e8e\u5b89\u5168\u8d44\u6599\u5185\u5bb9\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752715",
                "content": "\u4ea7\u54c1\u5b89\u5168\u7279\u6027\u63cf\u8ff0",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752716",
                "content": "\u4ea7\u54c1\u9632\u75c5\u6bd2\u3001\u52a0\u56fa\u6307\u5357",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752717",
                "content": "\u4ea7\u54c1\u5b89\u88c5\u6307\u5357",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752718",
                "content": "\u4ea7\u54c1\u5b89\u5168\u7ef4\u62a4\u624b\u518c",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752715,65752716",
        "answer": "65752715,65752716,65752718",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 51.1,
        "type": 1,
        "isanswerright": false,
        "ordernum": 34,
        "programitems": {}
    },
    {
        "id": 18359033,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4ee5\u4e0b\u54ea\u4e9b\u5c5e\u4e8e\u76d1\u542c\u63a5\u53e3\u5b89\u5168\u8981\u6c42( &nbsp; &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752719",
                "content": "\u5728\u6ca1\u6709\u534e\u4e3a\u660e\u786e\u9700\u6c42\u7684\u60c5\u51b5\u4e0b\uff0c\u4e25\u7981\u5f00\u53d1\u5177\u6709\u76d1\u542c\u6027\u8d28\u7684\u529f\u80fd\u548c\u63a5\u53e3\uff0c\u65e0\u8bba\u8be5\u529f\u80fd\u548c\u63a5\u53e3\u662f\u5426\u8981\u9075\u5faa\u76f8\u5e94\u7684\u56fd\u5bb6\u6807\u51c6\u548c\u56fd\u9645\u6807\u51c6",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752720",
                "content": "\u5728\u534e\u4e3a\u5bf9\u5408\u6cd5\u76d1\u542c\u63a5\u53e3\u6709\u9700\u6c42\u7684\u60c5\u51b5\u4e0b\uff0c\u5408\u4f5c\u65b9\u9700\u6839\u636e\u534e\u4e3a\u63d0\u4f9b\u7684\u76d1\u542c\u529f\u80fd\u6216\u63a5\u53e3\u7684\u6587\u4ef6\u4e2d\u7684\u8981\u6c42\u5f00\u53d1",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752721",
                "content": "\u5728\u6b63\u5e38\u4e1a\u52a1\u6d41\u7a0b\u548c\u6807\u51c6\u534f\u8bae\u4e4b\u5916\uff0c\u7981\u6b62\u63d0\u4f9b\u91c7\u96c6\u6700\u7ec8\u7528\u6237\u539f\u59cb\u901a\u4fe1\u5185\u5bb9\uff08\u8bed\u97f3\u7c7b\u3001\u77ed\u4fe1/\u5f69\u4fe1\u7c7b\u3001\u4f20\u771f\u7c7b\u3001\u6570\u636e\u4e1a\u52a1\u7c7b\uff09\u7684\u529f\u80fd\uff0c\u5373\u4f7f\u51fa\u4e8e\u4fdd\u969c\u7f51\u7edc\u8fd0\u8425\u548c\u670d\u52a1\u76ee\u7684",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752722",
                "content": "\u7cfb\u7edf\u652f\u6301\u65e0\u6cd5\u4ece\u7528\u6237\u9762\u76f4\u63a5\u767b\u9646\u8fde\u63a5\u7ba1\u7406\u63a5\u53e3\uff08\u4e0d\u652f\u6301\u72ec\u7acb\u7684\u7ba1\u7406IP\u5730\u5740\u7684\u4ea7\u54c1\u9664\u5916\uff09",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752719,65752720,65752721",
        "answer": "65752719,65752720,65752721",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 60.2,
        "type": 1,
        "isanswerright": true,
        "ordernum": 35,
        "programitems": {}
    },
    {
        "id": 18359034,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u5bf9\u63d0\u4f9b\u5408\u6cd5\u76d1\u542c\u63a5\u53e3\u7684\u4ea7\u54c1\u7248\u672c\u7684\u8981\u6c42\uff0c\u4ee5\u4e0b\u8bf4\u6cd5\u6b63\u786e\u7684\u662f\uff1f( &nbsp; &nbsp; &nbsp; &nbsp;)",
        "answeritems": [
            {
                "sn": "65752723",
                "content": "\u4ea7\u54c1\u63d0\u4f9b\u4e24\u4e2a\u7248\u672c\u7684\u8f6f\u4ef6\u5b89\u88c5\u5305\uff1a\u4e00\u4e2a\u652f\u6301\u5408\u6cd5\u76d1\u542c\uff0c\u4e00\u4e2a\u4e0d\u652f\u6301\u5408\u6cd5\u76d1\u542c\u3002\u6839\u636e\u5e02\u573a\u7684\u5b89\u5168\u8981\u6c42\uff0c\u9009\u62e9\u5bf9\u5e94\u7684\u8f6f\u4ef6\u5b89\u88c5\u5305\u8fdb\u884c\u90e8\u7f72",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752724",
                "content": "\u4ea7\u54c1\u63d0\u4f9b\u8f6f\u4ef6\u5b89\u88c5\u5305\u62c6\u5206\u4e3a\uff1a\u57fa\u672c\u8f6f\u4ef6\u5b89\u88c5\u5305\u548c\u5408\u6cd5\u76d1\u542c\u63d2\u4ef6\u5b89\u88c5\u5305\u3002\u6839\u636e\u5e02\u573a\u7684\u5b89\u5168\u8981\u6c42\uff0c\u9009\u62e9\u662f\u5426\u5b89\u88c5\u5408\u6cd5\u76d1\u542c\u63d2\u4ef6\u5b89\u88c5\u5305",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752725",
                "content": "\u4ea7\u54c1\u53ef\u4ee5\u901a\u8fc7License\u63a7\u5236\u5e26\u5408\u6cd5\u76d1\u542c\u63a5\u53e3\u7684\u7248\u672c\u5728\u67d0\u4e2a\u5730\u533a\u6216\u56fd\u5bb6\u9500\u552e\u6216\u4f7f\u7528",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752726",
                "content": "\u5fc5\u987b\u901a\u8fc7\u7248\u672c\u9694\u79bb\u786e\u4fdd\u7248\u672c\u4e2d\u53ea\u5b58\u5728\u7b26\u5408\u5f53\u5730\u5408\u6cd5\u76d1\u542c\u63a5\u53e3\u6807\u51c6\u7684\u4ee3\u7801",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752724,65752725,65752726",
        "answer": "65752723,65752724,65752726",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 600.0,
        "type": 1,
        "isanswerright": false,
        "ordernum": 36,
        "programitems": {}
    },
    {
        "id": 18359035,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u5728\u53e3\u4ee4\u5b89\u5168\u8981\u6c42\u4e2d\uff0c\u7cfb\u7edf\u53ef\u4ee5\u901a\u8fc7\u5982\u4e0b\u90a3\u51e0\u79cd\u65b9\u5f0f\u63d0\u4f9b\u89e3\u9501\u7528\u6237\u7684\u673a\u5236( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752727",
                "content": "\u5bf9\u4e8e\u53e3\u4ee4\u5c1d\u8bd5N\u6b21\u5931\u8d25\u88ab\u9501\u5b9a\u7684\u7528\u6237\uff0c\u7cfb\u7edf\u8981\u80fd\u591f\u8bbe\u7f6e\u81ea\u52a8\u89e3\u9501\u65f6\u95f4\u3002",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752728",
                "content": "\u7528\u6237\u88ab\u9501\u65f6\u95f4\u8fbe\u5230\u9884\u5b9a\u4e49\u65f6\u95f4\uff0c\u53ef\u81ea\u52a8\u89e3\u9501\u8be5\u7528\u6237\uff0c\u6216\u8005\u4e5f\u53ef\u901a\u8fc7\u5b89\u5168\u7ba1\u7406\u5458\u624b\u5de5\u89e3\u9501\u8be5\u7528\u6237",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752729",
                "content": "\u5728\u9501\u5b9a\u65f6\u95f4\u5185\uff0c\u4ec5\u80fd\u5141\u8bb8\u5e94\u7528\u5b89\u5168\u7ba1\u7406\u5458\u89d2\u8272\u6240\u5c5e\u5e10\u53f7\u624b\u52a8\u89e3\u9501\u8be5\u7528\u6237",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752730",
                "content": "\u88ab\u9501\u7528\u6237\u901a\u8fc7\u8fdb\u884c\u5bc6\u7801\u4fee\u6539\u6765\u8fbe\u5230\u81ea\u52a8\u89e3\u9501",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752727,65752728,65752729",
        "answer": "65752727,65752728,65752729",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 146.2,
        "type": 1,
        "isanswerright": true,
        "ordernum": 37,
        "programitems": {}
    },
    {
        "id": 18359036,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u5173\u4e8e\u53e3\u4ee4\u5b89\u5168\uff0c\u63cf\u8ff0\u6b63\u786e\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752731",
                "content": "\u53e3\u4ee4\u4e0d\u80fd\u5728\u7f51\u7edc\u4e2d\u660e\u6587\u4f20\u8f93\uff0c\u53e3\u4ee4\u7b49\u8ba4\u8bc1\u51ed\u8bc1\u5728\u4f20\u8f93\u8fc7\u7a0b\u4e2d\u5fc5\u987b\u52a0\u5bc6\uff0c\u4f7f\u7528\u9ad8\u5b89\u5168\u7b49\u7ea7\u7684\u52a0\u5bc6\u7b97\u6cd5",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752732",
                "content": "\u7528\u6237\u53ef\u4fee\u6539\u81ea\u5df1\u7684\u53e3\u4ee4\uff0c\u9700\u6ee1\u8db3\u5982\u4e0b\u8981\u6c42\uff1a1) \u7528\u6237\u4fee\u6539\u81ea\u5df1\u53e3\u4ee4\u65f6\u5fc5\u987b\u9a8c\u8bc1\u65e7\u53e3\u4ee4\uff1b2) \u4e0d\u5141\u8bb8\u4fee\u6539\u9664\u81ea\u8eab\u5e10\u53f7\u4ee5\u5916\u7684\u5e10\u53f7\u7684\u53e3\u4ee4\uff08\u7ba1\u7406\u5458\u9664\u5916\uff09",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752733",
                "content": "\u64cd\u4f5c\u754c\u9762\u4e2d\u7684\u53e3\u4ee4\u4e0d\u80fd\u660e\u6587\u663e\u793a",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752734",
                "content": "\u53e3\u4ee4\u8f93\u5165\u6846\u4e0d\u652f\u6301\u62f7\u8d1d\u529f\u80fd",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752731,65752732,65752733,65752734",
        "answer": "65752731,65752732,65752733,65752734",
        "score": 2.0,
        "userscore": 2.0,
        "answerexplain": "",
        "worktime": 227.2,
        "type": 1,
        "isanswerright": true,
        "ordernum": 38,
        "programitems": {}
    },
    {
        "id": 18359037,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u4e3a\u4e86\u4fdd\u8bc1\u64cd\u4f5c\u7cfb\u7edf\u7684\u5b89\u5168\uff0c\u9664\u4e86\u8bbe\u7f6e\u7528\u6237\u53e3\u4ee4\u548c\u63a7\u5236\u6587\u4ef6\u7684\u4f7f\u7528\u6743\u9650\uff0c\u8fd8\u9700\u8981\u91c7\u53d6\u54ea\u4e9b\u63aa\u65bd\uff1f ( &nbsp; &nbsp; )\uff1a",
        "answeritems": [
            {
                "sn": "65752735",
                "content": "\u5173\u952e\u8bbe\u5907\u7684\u9694\u79bb",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752736",
                "content": "\u8bbe\u5907\u7684\u5b8c\u6574\u6027",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752737",
                "content": "\u5b9a\u671f\u68c0\u67e5\u5b89\u5168\u65e5\u5fd7\u548c\u7cfb\u7edf\u72b6\u6001",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752738",
                "content": "\u9632\u75c5\u6bd2",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752735,65752736,65752737,65752738",
        "answer": "65752737,65752738",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 19.4,
        "type": 1,
        "isanswerright": false,
        "ordernum": 39,
        "programitems": {}
    },
    {
        "id": 18359038,
        "user": "icollege-migrate",
        "way": 1,
        "question": "\u5728\u7f51\u7edc\u5b89\u5168\u603b\u4f53\u8981\u6c42\u4e2d\uff0c\u4e0b\u9762\u5c5e\u4e8eB\u7c7b\u5b89\u5168\u8981\u6c42\u7684\u662f( &nbsp; &nbsp; &nbsp; &nbsp;)\uff1a",
        "answeritems": [
            {
                "sn": "65752739",
                "content": "\u8f6f\u4ef6\u5b8c\u6574\u6027\u4fdd\u62a4",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752740",
                "content": "\u534f\u8bae\u9632\u653b\u51fb",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752741",
                "content": "\u64cd\u4f5c\u7cfb\u7edf\u52a0\u56fa\u4e0e\u9632\u75c5\u6bd2",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752742",
                "content": "web\u5e94\u7528\u5b89\u5168",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752739,65752741",
        "answer": "65752740,65752741,65752742",
        "score": 2.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 123.4,
        "type": 1,
        "isanswerright": false,
        "ordernum": 40,
        "programitems": {}
    },
    {
        "id": 18358979,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u534e\u4e3a\u7f51\u7edc\u5b89\u5168\u7ea2\u7ebf\u603b\u4f53\u8981\u6c42\u5206A\u7c7b\u53caB\u7c7b\uff0c\u5176\u4e2dA\u7c7b\u5c5e\u4e8e\u4ea7\u54c1\u6838\u5fc3\u529f\u80fd\uff0c\u5982\u679c\u6ca1\u6709\u6b64\u529f\u80fd\u4ea7\u54c1\u4e0d\u53ef\u7528\u6216\u5b58\u5728\u91cd\u5927\u5b89\u5168\u9690\u60a3\u3002",
        "answeritems": [
            {
                "sn": "65752543",
                "content": "\u6b63\u786e",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752544",
                "content": "\u9519\u8bef",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752543",
        "answer": "65752543",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 22.4,
        "type": 2,
        "isanswerright": true,
        "ordernum": 41,
        "programitems": {}
    },
    {
        "id": 18358981,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u5bf9\u4e8e\u5408\u4f5c\u65b9\u7684\u4ea7\u54c1\u8981\u6c42\uff0c\u65b0\u7acb\u9879\u7684\u4ea7\u54c1\u7248\u672c\u53ef\u4ee5\u4f18\u5148\u6ee1\u8db3\u7f51\u7edc\u5b89\u5168\u603b\u4f53\u8981\u6c42\u4e2dA\u7c7b\u8981\u6c42\uff0cB\u7c7b\u8981\u6c42\u53ef\u4ee5\u6682\u65f6\u4e0d\u6ee1\u8db3\uff1b &nbsp;",
        "answeritems": [
            {
                "sn": "65752549",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752550",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752550",
        "answer": "65752550",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 8.0,
        "type": 2,
        "isanswerright": true,
        "ordernum": 42,
        "programitems": {}
    },
    {
        "id": 18358982,
        "user": "icollege-migrate",
        "way": 2,
        "question": "Web\u5b89\u5168\u4e2d\uff0c\u9a8c\u8bc1\u7801\u53ef\u4ee5\u591a\u6b21\u4f7f\u7528\uff0c\u65b0\u7684\u8fde\u63a5\u53ef\u4ee5\u4e0d\u9700\u8981\u91cd\u65b0\u751f\u6210\u65b0\u7684\u9a8c\u8bc1\u7801\uff1b",
        "answeritems": [
            {
                "sn": "65752551",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752552",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752552",
        "answer": "65752552",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 59.7,
        "type": 2,
        "isanswerright": true,
        "ordernum": 43,
        "programitems": {}
    },
    {
        "id": 18358983,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u5728\u975e\u5bf9\u79f0\u5bc6\u94a5\u5bc6\u7801\u4f53\u5236\u4e2d\uff0c\u53d1\u4fe1\u65b9\u4e0e\u6536\u4fe1\u65b9\u4f7f\u7528\u4e0d\u540c\u7684\u5bc6\u94a5\uff1b",
        "answeritems": [
            {
                "sn": "65752553",
                "content": "\u6b63\u786e",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752554",
                "content": "\u9519\u8bef",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752553",
        "answer": "65752553",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 8.2,
        "type": 2,
        "isanswerright": true,
        "ordernum": 44,
        "programitems": {}
    },
    {
        "id": 18358984,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u5bf9\u4e8e\u6d89\u53ca\u4ea7\u54c1\u77e5\u8bc6\u4ea7\u6743\u3001\u9ad8\u5371\u64cd\u4f5c\u3001\u53ef\u5916\u90e8\u8c03\u7528\u7684\u5185\u90e8\u63a5\u53e3\u7b49\u4e0d\u671f\u671b\u5411\u6240\u6709\u5ba2\u6237\u4eba\u5458\u516c\u5f00\u7684\u5185\u5bb9\uff0c\u4e0d\u80fd\u901a\u8fc7\u4efb\u4f55\u65b9\u5f0f\u5411\u5ba2\u6237\u516c\u5f00\uff1b&nbsp;",
        "answeritems": [
            {
                "sn": "65752555",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752556",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752556",
        "answer": "65752556",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 58.5,
        "type": 2,
        "isanswerright": true,
        "ordernum": 45,
        "programitems": {}
    },
    {
        "id": 18358985,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u8ba1\u7b97\u673a\u7cfb\u7edf\u7684\u8106\u5f31\u6027\u4e3b\u8981\u6765\u81ea\u4e8e\u7f51\u7edc\u64cd\u4f5c\u7cfb\u7edf\u7684\u4e0d\u5b89\u5168\u6027\uff1b&nbsp;",
        "answeritems": [
            {
                "sn": "65752557",
                "content": "\u6b63\u786e",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752558",
                "content": "\u9519\u8bef",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752557",
        "answer": "65752557",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 16.9,
        "type": 2,
        "isanswerright": true,
        "ordernum": 46,
        "programitems": {}
    },
    {
        "id": 18358986,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u64cd\u4f5c\u7cfb\u7edf\u4e2d\u8d85\u7ea7\u7528\u6237\u548c\u666e\u901a\u7528\u6237\u7684\u8bbf\u95ee\u6743\u9650\u6ca1\u6709\u5dee\u522b\uff1b",
        "answeritems": [
            {
                "sn": "65752559",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752560",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752560",
        "answer": "65752560",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 12.4,
        "type": 2,
        "isanswerright": true,
        "ordernum": 47,
        "programitems": {}
    },
    {
        "id": 18358987,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u4fdd\u62a4\u5e10\u6237\u3001\u53e3\u4ee4\u548c\u63a7\u5236\u8bbf\u95ee\u6743\u9650\u53ef\u4ee5\u63d0\u9ad8\u64cd\u4f5c\u7cfb\u7edf\u7684\u5b89\u5168\u6027\uff1b&nbsp;",
        "answeritems": [
            {
                "sn": "65752561",
                "content": "\u6b63\u786e",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752562",
                "content": "\u9519\u8bef",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752561",
        "answer": "65752561",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 8.2,
        "type": 2,
        "isanswerright": true,
        "ordernum": 48,
        "programitems": {}
    },
    {
        "id": 18358988,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u7ba1\u7406\u901a\u9053\u5b89\u5168\u4e3b\u8981\u662f\u7cfb\u7edf\u652f\u6301\u5bf9\u7ba1\u7406\u5e73\u9762\u7684\u63a5\u53e3\u7684\u8bbf\u95ee\u4fdd\u62a4\uff1b",
        "answeritems": [
            {
                "sn": "65752563",
                "content": "\u6b63\u786e",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752564",
                "content": "\u9519\u8bef",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752563",
        "answer": "65752563",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 46.0,
        "type": 2,
        "isanswerright": true,
        "ordernum": 49,
        "programitems": {}
    },
    {
        "id": 18358989,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u64cd\u4f5c\u7cfb\u7edf\u5b89\u5168\u53ea\u8981\u6c42\u5bf9\u7cfb\u7edf\u8fdb\u884c\u9632\u75c5\u6bd2\u5904\u7406\uff1b",
        "answeritems": [
            {
                "sn": "65752565",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752566",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752566",
        "answer": "65752566",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 4.5,
        "type": 2,
        "isanswerright": true,
        "ordernum": 50,
        "programitems": {}
    },
    {
        "id": 18358990,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u5b9a\u671f\u68c0\u67e5\u64cd\u4f5c\u7cfb\u7edf\u7684\u5b89\u5168\u65e5\u5fd7\u548c\u7cfb\u7edf\u72b6\u6001\u53ef\u4ee5\u6709\u52a9\u4e8e\u63d0\u9ad8\u64cd\u4f5c\u7cfb\u7edf\u5b89\u5168\uff1b",
        "answeritems": [
            {
                "sn": "65752567",
                "content": "\u6b63\u786e",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752568",
                "content": "\u9519\u8bef",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752567",
        "answer": "65752567",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 3.4,
        "type": 2,
        "isanswerright": true,
        "ordernum": 51,
        "programitems": {}
    },
    {
        "id": 18358991,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u5bf9\u64cd\u4f5c\u7cfb\u7edf\u7684\u9632\u75c5\u6bd2\u5904\u7406\u4e2d\uff0c\u53ea\u8981\u5408\u4f5c\u65b9\u4f7f\u7528\u4e3b\u6d41\u9632\u75c5\u6bd2\u8f6f\u4ef6\uff0c\u534e\u4e3a\u516c\u53f8\u53ef\u4ee5\u4ee5\u5408\u4f5c\u65b9\u7684\u626b\u63cf\u7ed3\u679c\u4e3a\u51c6\uff1b&nbsp;",
        "answeritems": [
            {
                "sn": "65752569",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752570",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752570",
        "answer": "65752570",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 18.2,
        "type": 2,
        "isanswerright": true,
        "ordernum": 52,
        "programitems": {}
    },
    {
        "id": 18358992,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u767b\u5f55\u8fc7\u7a0b\u4e2d\uff0c\u5f80\u670d\u52a1\u5668\u7aef\u4f20\u9012\u7528\u6237\u540d\u548c\u53e3\u4ee4\u65f6\uff0c\u4efb\u4f55\u573a\u666f\u4e0b\u90fd\u9700\u8981\u91c7\u7528HTTPS\u5b89\u5168\u534f\u8bae\uff1b &nbsp;",
        "answeritems": [
            {
                "sn": "65752571",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752572",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752572",
        "answer": "65752572",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 40.7,
        "type": 2,
        "isanswerright": true,
        "ordernum": 53,
        "programitems": {}
    },
    {
        "id": 18358993,
        "user": "icollege-migrate",
        "way": 2,
        "question": "Web\u5e94\u7528\u4e2d\uff0c\u5bf9\u7528\u6237\u7684\u8ba4\u8bc1\u53ef\u4ee5\u5728\u5ba2\u6237\u7aef\u8fdb\u884c\uff0c\u4e5f\u53ef\u4ee5\u5728\u670d\u52a1\u7aef\u96c6\u4e2d\u8fdb\u884c\uff1b",
        "answeritems": [
            {
                "sn": "65752573",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752574",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752574",
        "answer": "65752574",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 13.8,
        "type": 2,
        "isanswerright": true,
        "ordernum": 54,
        "programitems": {}
    },
    {
        "id": 18358994,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u7528\u4f5c\u5185\u90e8\u6d4b\u8bd5\u4f7f\u7528\u7684\u63a5\u53e3\uff0c\u53ea\u8981\u4e0d\u5728\u8d44\u6599\u4e2d\u516c\u5f00\uff0c\u53ef\u4ee5\u5728\u4ea7\u54c1\u6b63\u5f0f\u53d1\u5e03\u4e2d\u4fdd\u7559\uff1b",
        "answeritems": [
            {
                "sn": "65752575",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752576",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752576",
        "answer": "65752576",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 4.9,
        "type": 2,
        "isanswerright": true,
        "ordernum": 55,
        "programitems": {}
    },
    {
        "id": 18358995,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u53ea\u8981\u8fdb\u884c\u4e25\u683c\u6d4b\u8bd5\u4e14\u6709\u5b8c\u6574\u7684\u6d4b\u8bd5\u6570\u636e\u53ca\u62a5\u544a\uff0c\u4ea7\u54c1\u4e2d\u53ef\u4ee5\u4f7f\u7528\u79c1\u6709\u52a0\u5bc6\u7b97\u6cd5\uff1b",
        "answeritems": [
            {
                "sn": "65752577",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752578",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752578",
        "answer": "65752578",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 23.2,
        "type": 2,
        "isanswerright": true,
        "ordernum": 56,
        "programitems": {}
    },
    {
        "id": 18358996,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u7528\u4e8e\u654f\u611f\u6570\u636e\u4f20\u8f93\u52a0\u5bc6\u7684\u5bc6\u94a5\uff0c\u53ef\u4ee5\u786c\u7f16\u7801\u5728\u4ee3\u7801\u4e2d\uff1b&nbsp;",
        "answeritems": [
            {
                "sn": "65752579",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752580",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752580",
        "answer": "65752580",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 3.4,
        "type": 2,
        "isanswerright": true,
        "ordernum": 57,
        "programitems": {}
    },
    {
        "id": 18358997,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u53ea\u8981\u63a7\u5236\u65e5\u5fd7\u3001\u8bdd\u5355\u7b49\u6587\u4ef6\u7684\u8bbf\u95ee\u6743\u9650\uff0c\u53ef\u4ee5\u5728\u65e5\u5fd7\u3001\u8bdd\u5355\u7b49\u6587\u4ef6\u4e2d\u53ef\u4ee5\u8bb0\u5f55\u53e3\u4ee4\u3001\u94f6\u884c\u8d26\u53f7\u7b49\u654f\u611f\u6570\u636e\uff1b &nbsp;",
        "answeritems": [
            {
                "sn": "65752581",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752582",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752582",
        "answer": "65752582",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 8.4,
        "type": 2,
        "isanswerright": true,
        "ordernum": 58,
        "programitems": {}
    },
    {
        "id": 18358998,
        "user": "icollege-migrate",
        "way": 2,
        "question": "GDPR\uff08General Data Protection Regulation\uff09\u662f\u6b27\u76df\u5236\u5b9a\u7edf\u4e00\u6570\u636e\u4fdd\u62a4\u6cd5\uff0c\u4e0d\u9075\u5b88GDPR\uff0c\u5c06\u5bfc\u81f4\u7f5a\u6b3e\uff1a\u9ad8\u8fbe2000\u4e07\u6b27\u5143\u6216\u516c\u53f8\u5168\u7403\u603b\u8425\u4e1a\u989d4%\uff08\u4ee5\u8f83\u5927\u8005\u4e3a\u51c6\uff09\u3002",
        "answeritems": [
            {
                "sn": "65752583",
                "content": "\u6b63\u786e",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752584",
                "content": "\u9519\u8bef",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752583",
        "answer": "65752583",
        "score": 1.0,
        "userscore": 1.0,
        "answerexplain": "",
        "worktime": 18.7,
        "type": 2,
        "isanswerright": true,
        "ordernum": 59,
        "programitems": {}
    },
    {
        "id": 18358999,
        "user": "icollege-migrate",
        "way": 2,
        "question": "\u53ea\u8981\u9075\u5faa\u56fd\u9645\u6807\u51c6\u53ca\u6240\u5728\u56fd\u7684\u6cd5\u5f8b\u8981\u6c42\uff0c\u5408\u4f5c\u65b9\u53ef\u4ee5\u5728\u4ea7\u54c1\u4e2d\u63d0\u4f9b\u5408\u6cd5\u76d1\u542c\u63a5\u53e3\u53ca\u80fd\u529b\uff1b&nbsp;",
        "answeritems": [
            {
                "sn": "65752585",
                "content": "\u6b63\u786e",
                "isanswer": false,
                "score": 0.0,
                "count": 0,
                "percent": 0
            },
            {
                "sn": "65752586",
                "content": "\u9519\u8bef",
                "isanswer": true,
                "score": 0.0,
                "count": 0,
                "percent": 0
            }
        ],
        "useranswer": "65752585",
        "answer": "65752586",
        "score": 1.0,
        "userscore": 0.0,
        "answerexplain": "",
        "worktime": 0.0,
        "type": 2,
        "isanswerright": false,
        "ordernum": 60,
        "programitems": {}
    }
]`

type Data struct {
	id     int
	answer string
}

func getAnswers(s string) {
	ds := make([]Data, 0)
	err := json.Unmarshal([]byte(s), &ds)
	if err != nil {
		log.Println(err)
		return
	}
	answer := make(map[int]string)
	for _, a := range ds {
		answer[a.id] = a.answer
	}
}
func RunBrowser() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	b, err := chromedp.NewBrowser(taskCtx, "www.baidu.com")
	if err != nil {
		log.Println(err)
		return
	}
	//b.Execute(taskCtx,)
	_ = b
	err = chromedp.Run(
		taskCtx,
		visit(),
	)
	if err != nil {
		log.Println(err)
	}
}
func visit() chromedp.Tasks {
	var buf []byte
	log.Println("begin...")
	return chromedp.Tasks{
		chromedp.ActionFunc(func(context.Context) error {
			log.Println("tasks")
			return nil
		}),

		chromedp.Navigate("https://github.com/chromedp/chromedp"),
		chromedp.ActionFunc(func(context.Context) error {
			log.Println("open web")
			return nil
		}),
		chromedp.WaitVisible(`.mb-3 h4`, chromedp.ByQuery),
		chromedp.ActionFunc(func(context.Context) error {
			log.Println("find")
			return nil
		}),
		chromedp.Screenshot(`#post`, &buf, chromedp.ByQuery, chromedp.NodeVisible),
		chromedp.ActionFunc(func(context.Context) error {
			log.Println("shot")
			return nil
		}),
		chromedp.ActionFunc(func(context.Context) error {
			return ioutil.WriteFile("./mojotv_local.png", buf, 0644)
		}),
		chromedp.ActionFunc(func(context.Context) error {
			log.Println("save")
			return nil
		}),
	}
}
func TestRunBrowser(t *testing.T) {
	RunBrowser()
}
