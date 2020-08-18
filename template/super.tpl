{{template "header" .}}
<div class="row">
    <div class="col">

        <table class="table table-hover">
            <thead>
            <tr>
                <th scope="col">项目</th>
                <th scope="col">状态</th>
            </tr>
            </thead>
            <tbody>

            {{range $index, $element := .data}}
                <tr>
                    <td>
                        {{ $index }}
                    </td>
                    <td>
                        {{if eq $element "RUNNING"}}
                            <span class="badge badge-pill badge-success">{{$element}}</span>
                        {{ else if eq $element "STARTING" }}
                            <span class="badge badge-pill badge-primary">{{$element}}</span>
                        {{ else if eq $element "FATAL" }}
                            <span class="badge badge-pill badge-danger">{{$element}}</span>
                        {{ else if eq $element "STOP" }}
                            <span class="badge badge-pill badge-warning">{{$element}}</span>
                        {{else}}
                            <span class="badge badge-pill badge-info">{{$element}}</span>
                        {{end}}
                    </td>

                </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>
{{template "footer" .}}

