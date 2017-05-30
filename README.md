# gogo
Go语言


1、git init //初始化项目，执行完此命令后会生成一个.git文件夹

2、git add . //将本地项目所有文件添加到git管理，.指全部文件

3、git commit -m “提交描述"

4、git remote add origin 刚刚新建的Github地址 //将本地项目与远程git仓库关联

5、git push -u origin master //执行此命令如果出现错误，应该是README.md文件在本地项目中不存在从而导致冲突，我的一贯解决办法就是用这个命令git push -f origin master，强制将本地项目push到远程仓库。在平常的操作中，用这个强制的命令很可能会出现很多问题，建议不要用，不过此处是初始化项目，实用这个命令就不会有什么问题了。 