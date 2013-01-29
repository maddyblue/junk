package appstats

const HTML_BASE = `
{{ define "top" }}<!DOCTYPE html>
<html lang="en">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
  <style>
    @import "static/appstats_css.css";
  </style>
  <title>Appstats - {{.Env.APPLICATION_ID}}</title>
{{ end }}

{{ define "body" }}
</head>
<body>
  <div class="g-doc">
    {{/* Header begin */}}
    <div id="hd" class="g-section">
      <div class="g-section">
        <a href="."><img id="ae-logo" src="static/app_engine_logo_sm.gif"
          width="176" height="30" alt="Google App Engine" border="0"></a>
      </div>
      <div id="ae-appbar-lrg" class="g-section">
        <div class="g-section g-tpl-50-50 g-split">
          <div class="g-unit g-first">
            <h1>Application Stats for {{.Env.APPLICATION_ID}}</h1>
          </div>
          <div class="g-unit">
            All costs displayed in micropennies (1 dollar equals 100 pennies, 1 penny equals 1 million micropennies)
          </div>
        </div>
      </div>
    </div>
    {{/* Header end */}}
    {{/* Body begin */}}
    <div id="bd" class="g-section">
      {{/* Content begin */}}
      <div>
{{ end }}

{{ define "end" }}
      </div>
      {{/* Content end */}}
    </div>
    {{/* Body end */}}
  </div>
<script src="static/appstats_js.js"></script>
{{ end }}

{{ define "footer" }}
</body>
</html>
{{ end }}
`

const HTML_MAIN = `
{{ define "main" }}
{{ template "top" . }}
{{ template "body" . }}

<form id="ae-stats-refresh" action=".">
  <button id="ae-refresh">Refresh Now</button>
</form>

{{ if .Requests }}
<div class="g-section g-tpl-33-67">
  <div class="g-unit g-first">
    {{/* RPC stats table begin */}}
    <div class="ae-table-wrapper-left">
      <div class="ae-table-title">
        <div class="g-section g-tpl-50-50 g-split">
          <div class="g-unit g-first"><h2>RPC Stats</h2></div>
          <div id="ae-rpc-expand-all" class="g-unit"></div>
        </div>
      </div>
      <table cellspacing="0" cellpadding="0" class="ae-table ae-stripe" id="ae-table-rpc">
        <colgroup>
          <col id="ae-rpc-label-col">
          <col id="ae-rpc-stats-col">
        </colgroup>
        <thead>
          <tr>
            <th>RPC</th>
            <th>Count</th>
            <th>Cost</th>
            <th>Cost&nbsp;%</th>
          </tr>
        </thead>
        {{ range $index, $item := .AllStatsByCount }}
        <tbody>
          <tr>
            <td>
              <span class="goog-inline-block ae-zippy ae-zippy-expand" id="ae-rpc-expand-{{$index}}"></span>
              {{$item.Name}}
            </td>
            <td>{{$item.Count}}</td>
            <td title="">{{$item.Cost}}</td>
            <td>{{$item.CostPct}}</td>
          </tr>
        </tbody>
        <tbody class="ae-rpc-detail" id="ae-rpc-expand-{{$index}}-detail">
          {{ range $subitem := $item.SubStats }}
          <tr>
            <td class="rpc-req">{{$subitem.Name}}</td>
            <td>{{$subitem.Count}}</td>
            <td title="">{{$subitem.Cost}}</td>
            <td>{{$subitem.CostPct}}</td>
          </tr>
          {{ end }}
        </tbody>
        {{ end }}
      </table>
    </div>
    {{/* RPC stats table end */}}
  </div>
  <div class="g-unit">
    {{/* Path stats table begin */}}
    <div class="ae-table-wrapper-right">
      <div class="ae-table-title">
        <div class="g-section g-tpl-50-50 g-split">
          <div class="g-unit g-first"><h2>Path Stats</h2></div>
          <div class="g-unit" id="ae-path-expand-all"></div>
        </div>
      </div>
      <table cellspacing="0" cellpadding="0" class="ae-table" id="ae-table-path">
        <colgroup>
          <col id="ae-path-label-col">
          <col id="ae-path-rpcs-col">
          <col id="ae-path-reqs-col">
          <col id="ae-path-stats-col">
        </colgroup>
        <thead>
          <tr>
            <th>Path</th>
            <th>#RPCs</th>
            <th>Cost</th>
            <th>Cost%</th>
            <th>#Requests</th>
            <th>Most Recent requests</th>
          </tr>
        </thead>
        {{ range $index, $item := .PathStatsByCount }}
        <tr>
          <td>
            <span class="goog-inline-block ae-zippy ae-zippy-expand" id="ae-path-expand-{{$index}}"></span>
            {{$item.Name}}
          </td>
          <td>
            {{$item.Count}}
          </td>
          <td title="">{{$item.Cost}}</td>
          <td>{{$item.CostPct}}%</td>
          <td>{{$item.Requests}}</td>
          <td>
            {{ range $index := $item.RecentReqs }}
                {{ if $index }} <a href="#req-{{$index}}">({{$index}})</a> {{ else }} ... {{ end }}
            {{ end }}
          </td>
          <tbody class="path path-{{$index}}">
            {{ range $subitem := $item.SubStats }}
            <tr>
              <td class="rpc-req">{{$subitem.Name}}</td>
              <td>{{$subitem.Count}}</td>
              <td title="">{{$subitem.Cost}}</td>
              <td>{{$subitem.CostPct}}%</td>
              <td></td>
              <td></td>
            </tr>
            {{ end }}
          </tbody>
        {{ end }}
      </table>
    </div>
    {{/* Path stats table end */}}
  </div>
</div>
<div id="ae-req-history">
  <div class="ae-table-title">
    <div class="g-section g-tpl-50-50 g-split">
      <div class="g-unit g-first"><h2>Requests History</h2></div>
      <div class="g-unit" id="ae-request-expand-all"></div>
    </div>
  </div>

  <table cellspacing="0" cellpadding="0" class="ae-table" id='ae-table-request'>
    <colgroup>
      <col id="ae-reqs-label-col">
    </colgroup>
    <thead>
      <tr>
        <th colspan="4">Request</th>
      </tr>
    </thead>
    {{ range $index, $r := .Requests }}
    <tbody>
      <tr>
        <td colspan="4" class="ae-hanging-indent">
          <span class="goog-inline-block ae-zippy ae-zippy-expand" id="ae-path-requests-{{$index}}"></span>
          ({{$index}})
          <a name="req-{{$index}}" href="details?time={{$r.RequestStats.Start.Nanosecond}}" class="ae-stats-request-link">
            {{$r.RequestStats.Start}}
            "{{$r.RequestStats.Method}}
            {{$r.RequestStats.Path}}{{if $r.RequestStats.Query}}?{{$r.RequestStats.Query}}{{end}}"
            {{if $r.RequestStats.Status}}{{$r.RequestStats.Status}}{{end}}
          </a>
          real={{$r.RequestStats.Duration}}
          {{/*
          api={{$r.api_milliseconds}}
          overhead={{$r.overhead_walltime_milliseconds}}ms
          ({{$r.combined_rpc_count}} RPC{{$r.combined_rpc_count}},
            cost={{$r.combined_rpc_cost_micropennies}},
            billed_ops=[{{$r.combined_rpc_billed_ops}}])
          */}}
          ({{$r.RequestStats.RPCStats | len}} RPCs)
        </td>
      </tr>
    </tbody>
    <tbody class="reqon" id="ae-path-requests-{{$index}}-tbody">
      {{ range $item := $r.SubStats }}
      <tr>
        <td class="rpc-req">{{$item.Name}}</td>
        <td>{{$item.Count}}</td>

        <td>{{$item.Cost}}</td>
        {{/*<td>{{$item.total_billed_ops_str}}</td>*/}}
      </tr>
      {{ end }}
    </tbody>
    {{ end }}
  </table>
</div>
{{ else }}
<div>
  No requests have been recorded yet.  While it is possible that you
  simply need to wait until your server receives some requests, this
  is often caused by a configuration problem.
  <!-- TODO:maxr templatize python/java in the below link -->
  <a href="https://developers.google.com/appengine/docs/python/tools/appstats#EventRecorders"
  >Learn more</a>
</div>
{{ end }}

{{ template "end" . }}

<script>
  var z1 = new ae.Stats.MakeZippys('ae-table-rpc', 'ae-rpc-expand-all');
  var z2 = new ae.Stats.MakeZippys('ae-table-path', 'ae-path-expand-all');
  var z3 = new ae.Stats.MakeZippys('ae-table-request', 'ae-request-expand-all');
</script>

{{ template "footer" . }}
{{ end }}
`
