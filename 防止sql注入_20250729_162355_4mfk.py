// 代码生成时间: 2025-07-29 16:23:54
以下是用Django框架创建的应用组件，实现了防止SQL注入的功能。

```python
# models.py
from django.db import models

class Book(models.Model):
    """
    书籍模型
    """
    title = models.CharField(max_length=100, help_text="书籍标题")
    author = models.CharField(max_length=100, help_text="作者")
    publish_date = models.DateField(help_text="发布日期")

    def __str__(self):
        return self.title

# views.py
from django.shortcuts import render, get_object_or_404
from .models import Book

def list_books(request):
    """
    列出所有书籍的视图
    """
    try:
        books = Book.objects.all()  # 使用Django的查询方法防止SQL注入
        return render(request, 'books/list.html', {'books': books})
    except Exception as e:
        # 错误处理
        return render(request, 'books/error.html', {'error': str(e)})

def detail_book(request, id):
    """
    书籍详情视图
    """
    try:
        book = get_object_or_404(Book, pk=id)  # 同样使用Django的查询方法
        return render(request, 'books/detail.html', {'book': book})
    except Exception as e:
        # 错误处理
        return render(request, 'books/error.html', {'error': str(e)})

# urls.py
from django.urls import path
from . import views

urlpatterns = [
    path('', views.list_books, name='list-books'),
    path('book/<int:id>/', views.detail_book, name='detail-book'),
]
```

以上代码包含以下组成部分：

1. models.py：定义了书籍模型（防止SQL注入）。

2. views.py：
   - `list_books`视图：列出所有书籍，使用Django的查询方法防止SQL注入。
   - `detail_book`视图：根据ID获取书籍详情，同样使用Django的查询方法防止SQL注入。
   - 错误处理：捕获异常，返回错误页面。

3. urls.py：定义了路由，将URL映射到相应的视图函数。

只要遵循Django的最佳实践，使用Django的ORM查询方法，就可以很好地防止SQL注入攻击。

希望这些代码对你有所帮助！如果你有任何其他问题，欢迎随时问我。