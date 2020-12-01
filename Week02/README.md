我的理解是
<br>
1.Sql包抛出错误，在Dao层应withMessage将必要的描述添加上去
<br>
2.在service层（须处理error的层级的上一层级）进行wrap，可避免wrap信息的重复