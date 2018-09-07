# MN_Exercise_Webserver
The task was to create a RESTful API that returns a random password with a requested number of characters. I created it using AWS provider and terraform as my orchestration tool. 

## Installation and Usage:
You need to install terraform first and it's scripting can be done in notepad++. To create a server I used GO language, the editor I used for it is Intellij and both are easily available for download. You must have terraform files, webserver and key file that you are using in one folder.Please enter your key name in assignment.tf and place that key in the same folder as other files. Kinldy enter your AWS access and secret key in variables.tf file. Then you must run the following commands in order to run it

1. terraform.exe init
2. terraform plan -out usman-demo.plan (usman-demo is the server name here)
3. terraform apply webserver.plan 
  

## Why I used AWS:
My pick was between AWS and AZURE because I have worked with them being a Software Quality Assurance Engineer at Cloudplex. Both providers are really good and have neck to neck competition in today’s market. In dramatic words you can call AWS iis the Superman and AZURE is the Batman of today’s world because they are this much close these days. 
Reason for choosing AWS over Azure here was that, AWS has a larger global footprint and has a broad ecosystem of offerings in the AWS solution. Scalability is the key word here for AWS. In my personal experience I have done a lot of testing for both AWS and AZURE and I find AWS more user friendly, robust and is simply the universe of DevOps tools that is ever expanding.

## Why I picked Terraform:
There are so many configuration management tools like Terraform, puppet, ansible, Chef, Cloudformation etc. It was really hard going for one as they all have their advantages. First I got inclined towards Cloudformation, although it is a very good tool, but I had to drop it because it is not an open source tool and is AWS only. 
I preferred Terraform and Cloudformation because both were orchestration tools which means they both are designed to provision the servers themselves, leaving the job of configuring those servers to other tools and that is exactly what was required. I dropped Cloudformation for the already mentioned reason. Other tools like Chef, Puppet, Ansible and Saltstack are all designed to install and manage software on existing servers. Terraform also uses a declarative style of coding and I personally prefer that because the code specifies our desired end state which is always good.

## Challenges Faced:
As I have also mentioned in previous interview that I have worked in the test phase of the DevOps in which I did automation testing using selenium webdriver and API testing as far as my experience is concerned. So this was fairly a new thing to me. Because of my experience of working with different providers like AWS and AZURE, when I saw the assignment I got the sketch that exactly what I need to do. For implementation I had to read alot to learn about different configuration management tools and how can we use them. I had to learn scripting for Terraform and Go language. So this whole task was a challenge for me but as I am fascinated with DevOps and wish to have a career in this field, I enjoyed learning new things and doing it.

