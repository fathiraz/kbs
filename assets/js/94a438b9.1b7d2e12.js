"use strict";(self.webpackChunkdoc=self.webpackChunkdoc||[]).push([[739],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return d}});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=a.createContext({}),u=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=u(e.components);return a.createElement(s.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,s=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),m=u(n),d=r,k=m["".concat(s,".").concat(d)]||m[d]||c[d]||i;return n?a.createElement(k,o(o({ref:t},p),{},{components:n})):a.createElement(k,o({ref:t},p))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,o=new Array(i);o[0]=m;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:r,o[1]=l;for(var u=2;u<i;u++)o[u]=n[u];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},7113:function(e,t,n){n.r(t),n.d(t,{assets:function(){return p},contentTitle:function(){return s},default:function(){return d},frontMatter:function(){return l},metadata:function(){return u},toc:function(){return c}});var a=n(7462),r=n(3366),i=(n(7294),n(3905)),o=["components"],l={sidebar_label:"Answer",sidebar_position:1},s="Answer",u={unversionedId:"problem-5/answer",id:"problem-5/answer",title:"Answer",description:"Notification service is a service for sending notifier to user, including :",source:"@site/docs/problem-5/answer.md",sourceDirName:"problem-5",slug:"/problem-5/answer",permalink:"/kbs/docs/problem-5/answer",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/problem-5/answer.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_label:"Answer",sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Question",permalink:"/kbs/docs/problem-5/question"}},p={},c=[{value:"What you&#39;ll need",id:"what-youll-need",level:2},{value:"Prerequisites",id:"prerequisites",level:2},{value:"Default configuration",id:"default-configuration",level:2},{value:"Build the app",id:"build-the-app",level:2},{value:"Running the test",id:"running-the-test",level:2},{value:"Unit test",id:"unit-test",level:3},{value:"Unit test + coverage",id:"unit-test--coverage",level:3},{value:"Running the app",id:"running-the-app",level:2},{value:"Send Sms",id:"send-sms",level:3},{value:"Toggle Vendor",id:"toggle-vendor",level:3},{value:"Get All Vendor",id:"get-all-vendor",level:3},{value:"Source code",id:"source-code",level:2}],m={toc:c};function d(e){var t=e.components,n=(0,r.Z)(e,o);return(0,i.kt)("wrapper",(0,a.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"answer"},"Answer"),(0,i.kt)("p",{align:"center"},(0,i.kt)("img",{width:"150",height:"174",src:"https://raw.githubusercontent.com/fathiraz/kbs/main/ntf/logo.png"})),(0,i.kt)("h1",{id:"notification-service"},"Notification Service"),(0,i.kt)("p",null,"Notification service is a service for sending notifier to user, including :"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"sms"),(0,i.kt)("li",{parentName:"ul"},"email (TODO)"),(0,i.kt)("li",{parentName:"ul"},"push notification (TODO)")),(0,i.kt)("h2",{id:"what-youll-need"},"What you'll need"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://go.dev/dl//"},"Golang")," version 1.17 or above"),(0,i.kt)("li",{parentName:"ul"},"Basic knowledge about go mod ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/golang/go/wiki/Modules"},"https://github.com/golang/go/wiki/Modules"))),(0,i.kt)("h2",{id:"prerequisites"},"Prerequisites"),(0,i.kt)("p",null,"Download all required library using"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"make download\n")),(0,i.kt)("p",null,"Or you can use"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"go mod download\n")),(0,i.kt)("h2",{id:"default-configuration"},"Default configuration"),(0,i.kt)("p",null,"You can change default configuration on file .env"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-dotenv"},"# application config\nAPP_NAME=notification-service\nAPP_PORT=3000\nAPP_DEBUG=true\n\n# authentication config\nBASIC_USERNAME=kita\nBASIC_PASSWORD=bisa\n\n# sms vendor config\nSMS_CONFIG_TYPE=toml\nSMS_CONFIG_PATH=config/sms_vendor\n")),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"APP_NAME")," is configuration to set application name"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"APP_PORT")," is configuration to set application api port, accept number"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"APP_DEBUG")," is configuration to set application debug, accept only ",(0,i.kt)("strong",{parentName:"li"},"true")," / ",(0,i.kt)("strong",{parentName:"li"},"false")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"BASIC_USERNAME")," is configuration to set basic username on request header"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"BASIC_PASSWORD")," is configuration to set basic password on request header"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"SMS_CONFIG_TYPE")," is configuration to set sms config type to use on the application, accept only ",(0,i.kt)("strong",{parentName:"li"},"toml")," / ",(0,i.kt)("strong",{parentName:"li"},"yaml"),", you can check example of configuration each file in ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/fathiraz/kbs/tree/main/ntf/config/sms_vendor"},"here"),"."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"SMS_CONFIG_PATH")," is configuration to set sms config path to use on the application")),(0,i.kt)("h2",{id:"build-the-app"},"Build the app"),(0,i.kt)("p",null,"You can build the app using one of these method, choose whats you prefer"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"direct running"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"make run\n")),"  Or you can use",(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"go run main.go\n"))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"docker running"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"make run-docker\n")),"Or you can use",(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"docker-compose up\n")))),(0,i.kt)("h2",{id:"running-the-test"},"Running the test"),(0,i.kt)("p",null,"The app is testable and use colorized output test by ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/rakyll/gotest"},"gotest"),". This test should be run with file ",(0,i.kt)("inlineCode",{parentName:"p"},".env"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"config/sms_vendor/sms_vendor.yaml"),", & ",(0,i.kt)("inlineCode",{parentName:"p"},"config/sms_vendor/sms_vendor.toml")," has been set."),(0,i.kt)("h3",{id:"unit-test"},"Unit test"),(0,i.kt)("p",null,"Basic test with no coverage."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"make test\n")),(0,i.kt)("h3",{id:"unit-test--coverage"},"Unit test + coverage"),(0,i.kt)("p",null,"Test with coverage that will produced output to file ",(0,i.kt)("inlineCode",{parentName:"p"},"coverage-data.txt"),"."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"make cover\n")),(0,i.kt)("h2",{id:"running-the-app"},"Running the app"),(0,i.kt)("p",null,"You can build the app first on ",(0,i.kt)("a",{parentName:"p",href:"https://fathiraz.github.io/kbs/docs/problem-5/answer#build-the-app"},"this step"),". Import postman environment and collection ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/fathiraz/kbs/tree/main/ntf/doc"},"right here"),"."),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Postman collection ",(0,i.kt)("inlineCode",{parentName:"li"},"doc/Kitabisa.postman_collection.json")),(0,i.kt)("li",{parentName:"ul"},"Postman environment ",(0,i.kt)("inlineCode",{parentName:"li"},"doc/Kitabisa.\tpostman_environment.json"))),(0,i.kt)("p",null,"After running the app here is what you can do"),(0,i.kt)("h3",{id:"send-sms"},"Send Sms"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"method ",(0,i.kt)("inlineCode",{parentName:"li"},"POST")),(0,i.kt)("li",{parentName:"ul"},"url ",(0,i.kt)("inlineCode",{parentName:"li"},"{{url}}/v1/sms/send"))),(0,i.kt)("p",null,"request"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"curl --location --request POST '{{url}}/v1/sms/send' \\\n--header 'Authorization: Basic {{basic_authentication}}' \\\n--header 'Content-Type: application/json' \\\n--data-raw '{\n    \"number\" : \"628987654321\",\n    \"message\" : \"test\"\n}'\n")),(0,i.kt)("p",null,"success response"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'{\n    "success": true,\n    "code": 200,\n    "message": "message send successfully"\n}\n')),(0,i.kt)("p",null,"failed response"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'{\n    "success": false,\n    "code": 400,\n    "message": "code=400, message=Number wajib diisi",\n    "error": {\n        "message": "Number wajib diisi"\n    }\n}\n')),(0,i.kt)("h3",{id:"toggle-vendor"},"Toggle Vendor"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"method ",(0,i.kt)("inlineCode",{parentName:"li"},"POST")),(0,i.kt)("li",{parentName:"ul"},"url ",(0,i.kt)("inlineCode",{parentName:"li"},"{{url}}/v1/sms/toggle"))),(0,i.kt)("p",null,"request"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"curl --location --request POST '{{url}}/v1/sms/toggle' \\\n--header 'Authorization: Basic {{basic_authentication}}' \\\n--header 'Content-Type: application/json' \\\n--data-raw '{\n    \"name\" : \"sms_vendor_kita\",\n    \"toggle\" : true\n}'\n")),(0,i.kt)("p",null,"response"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'{\n    "success": true,\n    "code": 200,\n    "message": "sms vendor toggled successfully",\n    "data": {\n        "name": "sms_vendor_kita",\n        "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",\n        "endpoint": "/vendor_sms_kita",\n        "token": "Basic abc123def456",\n        "is_default": false,\n        "is_active": true,\n        "timeout": "30s"\n    }\n}\n')),(0,i.kt)("h3",{id:"get-all-vendor"},"Get All Vendor"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"method ",(0,i.kt)("inlineCode",{parentName:"li"},"GET")),(0,i.kt)("li",{parentName:"ul"},"url ",(0,i.kt)("inlineCode",{parentName:"li"},"{{url}}/v1/sms"))),(0,i.kt)("p",null,"request"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"curl --location --request GET '{{url}}/v1/sms' \\\n--header 'Authorization: Basic {{basic_authentication}}'\n")),(0,i.kt)("p",null,"response"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'{\n    "success": true,\n    "code": 200,\n    "message": "sms vendor get successfully",\n    "data": [\n        {\n            "name": "sms_vendor_kita",\n            "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",\n            "endpoint": "/vendor_sms_kita",\n            "token": "Basic abc123def456",\n            "is_default": false,\n            "is_active": true,\n            "timeout": "30s"\n        },\n        {\n            "name": "sms_vendor_bisa",\n            "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",\n            "endpoint": "/vendor_sms_bisa",\n            "token": "Token abc123def456",\n            "is_default": false,\n            "is_active": false,\n            "timeout": "30s"\n        },\n        {\n            "name": "sms_vendor_kitabisa",\n            "url": "https://626a3f7053916a0fbdf7c364.mockapi.io",\n            "endpoint": "/vendor_sms_kitabisa",\n            "token": "Auth abc123def456",\n            "is_default": true,\n            "is_active": true,\n            "timeout": "30s"\n        }\n    ]\n}\n')),(0,i.kt)("h2",{id:"source-code"},"Source code"),(0,i.kt)("p",null,"You can check the source code right here ",(0,i.kt)("strong",{parentName:"p"},(0,i.kt)("a",{parentName:"strong",href:"https://github.com/fathiraz/kbs/tree/main/ntf"},"https://github.com/fathiraz/kbs/tree/main/ntf"))))}d.isMDXComponent=!0}}]);