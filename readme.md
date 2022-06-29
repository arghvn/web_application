What is a web application and how is it different from a website?
Web Application:
Web application or web application is software that is used using a web browser and through local networks or the Internet. Basically, web applications are applications that can only be used through the web. More attention is paid to the performance and functionality of web applications, and their content is not so important and is usually not full of content. For example, Gmail is a popular web application.

website :
A web site is an Internet space that contains one or more pages. A set of interconnected web pages that have a home page and are located on a server is called a web page. Websites are designed to provide complete information about a person, organization, company, business, etc. to web users and focus specifically on content.

Websites focus on content, but web applications focus on performance. A website displays the same content to everyone, but a web application processes and displays information based on its interactions with the user. That is, the content displayed for each user is different from other users based on age, interest, gender, etc. In order to be able to view or use the information of a website, a user must be connected to the Internet, but this is not the case in web applications. Just once is enough to load the web application, then the user can access it offline.

The offline web application works with a feature called the Application Cache. The application repository can store all sections of a site offline and offline. Using this feature on the site, the server sends all JS, CSS, HTML files and images and other available information to the user's browser to be stored in the computer memory. The function of the web application repository is that when the user is online and visiting the site, the application repository is automatically updated and new files replace the old files.


this project goals :
static webpage for saling organic juices with sections sale , shop , members
working with templates
MVC framework
database access
unit testing

we use localhost:8000 for it

The main parts of the site include Google Map, a place to connect to the site, a place to personalize, a place to sell, a place for personal information of people, a place to introduce products, a place to display the information of site owners, all using cookies.

in time exit the terminal write exit

What is a resource server?
A resource server is a server for access-protected resources. It handles authenticated requests from an app that has an access token.
a resource server is set up to recieve requests from remote client machines HTTP or similar protocol.
it determines if the requested resource such as a file on the server hard drive exists or not.
 if it does the server sends the requested file back to the client.
 the resource server requirements :
listen on a TCP port for requests
logic to handle request (error or response )

for this project we should use net/http standard lybrary
there are three objects in this package , http.ListenAndserve that use for checking request or response , http.Hanlde use for handle http handler interface and http.HandleFunc use for specific http handler interface.

using localhost :
Website designers need to execute their code to design and build their site. For this purpose, they use a computer browser to execute html, css and javascript code, and a web server and interpreter to execute php code and other server-side languages ​​so that they can execute their entire project.
The first way to run server-side languages ​​seems to be to use an online server or host to run the project, but this method has many problems. For this reason, we use a local hosting space or local hosts instead of hosts.
Localhost is a local space on a personal computer that creates a space like a host or a real server for us.
 We have to pay to provide a standard service. It is usually not cost-effective to run and review projects that are in the pilot phase. Also, apart from the cost issue, we may sometimes not be able to access the Internet and the server.
Another point is that whenever we want to make any changes and edits to our files on the computer, we must also apply these changes to the main host, which is very time consuming. So it is better to first test our edits and errors on a virtual web server and then upload it to the host after finalizing the work.
Another point that can be mentioned is that our code may also have potential security issues and be severely damaged by being on the Internet. So to maintain information security, it is better to test them first on your personal computer and after finalizing and fixing the bugs, we put them on the host.


to beautify the appearanace of the site and a movie for demo using peach.blender.org site
 copy it to project folder and add mp4 in our code (not done)

 html template :
 Differences between dynamic and static pages in web design
 
In designing web pages, some pages have fixed content and some have variable content based on user behavior. Accordingly, web pages are divided into two categories: static (dynamic) and dynamic (dynamic).

Static page:
A static page is a page of a website that has consistent content for all users. These pages, often implemented in HTML, CSS, and JavaScript, often do not send a request to the server to display content, or send a single, fixed request to the server. The server also sends the same response to these pages in all situations.

For example, a page about us or a contact on the site has fixed content. Static pages do not mean that they have never been edited or changed. Rather, their content can be changed, but most cases do not need to be constantly changed, so they have a fixed content.

Dynamic page:
Dynamic pages are pages that have different content for users.
In dynamic pages, we see a variety of content based on the choice or behavior of users. Like the meteorological pages where we see separate content based on the choice of province and city.

An example of a dynamic web page is the shopping cart page. The content of this page is different for the user, based on his previous choices.
Dynamic pages can be divided into two categories. Pages that are dynamic with user-side code and pages that are server-side dynamic become a dynamic web page.

we use this funcions for templating , template.new , template.parse andtemplate.execute.

mvc pattern use for design
it contained model , controller and view module
model takes the response from database
view shows this response
controller passes response to the view

for examlpe mvc used for log in in this case controller check the info and send it to model , model sens a response to controller and ffinally controller sends feedback to view
view is connected to client and html
