{{template "header" .}}
<div class="row">
    <div class="col">
        <table class="table table-hover">
            <thead>
            <tr>
                <th scope="col">手机|邮箱</th>
                <th scope="col">验证码</th>
            </tr>
            </thead>
            <tbody>

            {{range $index, $element := .codes}}
                <tr>
                    <td>
                        {{$index}}
                    </td>
                    <td>
                        {{ $element }}
                    </td>
                </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>
{{template "footer" .}}

