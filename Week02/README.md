# 本周作业：
### 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

#### 答题详见代码

### 个人理解

dao层不需要Wrap这个error，应该是抛到service层之后，由service层的业务来决定是否需要Wrap这个sql.ErrNoRows。