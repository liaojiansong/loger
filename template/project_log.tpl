{{template "header" .}}
<div class="row">
    <div class="col">

        <table class="table table-hover">
            <thead>
            <tr>
                <th scope="col"></th>
                <th scope="col">等级</th>
                <th scope="col">时间</th>
                <th scope="col">请求参数</th>
                <th scope="col">调用栈</th>
            </tr>
            </thead>
            <tbody>

            {{range $index, $element := .Data}}
                <tr>
                    <td>
                        {{ $index }}
                    </td>
                    <td>
                        {{if eq $element.Level "INFO"}}
                            <span class="badge badge-pill badge-info">Info</span>
                        {{ else if eq $element.Level "WARNING" }}
                            <span class="badge badge-pill badge-warning">Warning</span>
                        {{else}}
                            <span class="badge badge-pill badge-danger">Danger</span>
                        {{end}}
                    </td>
                    <td>{{ $element.Time }}</td>
                    <td><small>{{ $element.Params }}</small></td>
                    <td>
{{/*                        {{$.Title}}*/}}
                        <ul>
                            {{range $element.Trace}}
                                <li class="list">
                                    <small>{{.}}</small>
                                </li>
                            {{end}}
                        </ul>

                    </td>

                </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>
{{template "footer" .}}

