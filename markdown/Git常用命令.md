# Git常用命令

## 1 命令

- 设置git账号和邮箱

  ```
   git config --global user.email "you@example.com"
   git config --global user.name "Your Name"
  ```

  

## 2 问题

1. [...io/gitea/cmd/hook.go:66 hookSetup()] [E] LFS server support needs at least Git v2.1.2

   原因：远程仓库有文件，需要pull回来再提交

   *该提示为本地版本与远程仓库不同，可使用 git pull --rebase origin master 命令将远程仓库的文件拉下来覆盖本地文件，再输入 git push -u origin master 提交。或输入 git push origin master -f 强行提交并覆盖远程仓库内文件。*

