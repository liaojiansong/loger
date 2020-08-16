{{define "header"}}
    <!doctype html>
    <html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" integrity="sha384-JcKb8q3iqJ61gNV9KGb8thSsNjpSL0n8PARn9HuZOnIxN0hoP+VmmDGMN5t9UJ0Z" crossorigin="anonymous">

        <title>运维工具</title>
    </head>
    <body>
    <div class="container">
        <div class="row" style="margin-top: 50px;margin-bottom: 50px">
            <div class="col">
                <nav class="navbar navbar-expand-lg navbar-light">
                    <a class="navbar-brand" href="/">运维平台</a>
                    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarTogglerDemo02" aria-controls="navbarTogglerDemo02" aria-expanded="false" aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon"></span>

                    </button>
                    <div class="collapse navbar-collapse" id="navbarTogglerDemo02">
                        <ul class="navbar-nav mr-auto mt-2 mt-lg-0">
                            <li class="nav-item active">
                                <a class="nav-link" href="/log/ieas?line=10">ieas日志</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="/log/ieas?line=10">brm日志</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="/log/ieas?line=10">php日志</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="/log/ieas?line=10">nginx日志</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="/log/ieas?line=10">supervisor</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="#">验证码查看</a>
                            </li>

                        </ul>
                    </div>
                </nav>
            </div>

        </div>

{{end}}