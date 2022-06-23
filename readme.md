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