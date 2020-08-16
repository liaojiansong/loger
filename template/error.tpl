{{template "header" .}}
<div class="row">
    <div class="col">
        <table class="table table-hover">
            <thead>
            <tr>
                <th scope="col">错误编号</th>
                <th scope="col">错误信息</th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <td>{{.code}}</td>
                <td>{{.msg}}</td>
            </tr>

            </tbody>
        </table>
    </div>
</div>
{{template "footer" .}}