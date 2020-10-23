{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command is basically the native report-command, but adds some back-end functionalites in order for the rest to work :)


    Recommended Trigger type and trigger: Regex; \A-r(eport)?u(ser)?(\s+|\z)

    Credit: ye olde boi#7325 U-ID:665243449405997066
*/}}

{{/*CONFIG AREA START*/}}

{{$reports := 750730537571975298}} {{/*The channel where your reports are logged into.*/}}
{{$reportDiscussion := 750099460314628176}} {{/*Your channel where users talk to staff*/}}

{{/*CONFIG AREA END*/}}


{{/*ACTUAL CODE DO NOT TOUCH UNLESS YOU KNOW WHAT YOU ARE DOING*/}}
{{if not (ge (len .CmdArgs) 2)}}
    ```{{.Cmd}} <User:Mention/ID> <Reason:Text>```
    Not enough arguments passed.
{{else}}
    {{dbSet 2000 "reportLog" $reports}}
    {{dbSet 2000 "reportDiscussion" $reportDiscussion}}
    {{$secret := adjective}}
    {{$s := execAdmin "log"}}
    {{$user := userArg (index .CmdArgs 0)}}
    {{$reason := joinStr " " (slice .CmdArgs 1)}}
    {{$reportGuide := (printf "\nDismiss report with ❌, take action with 🛡️, or request more background information with ⚠️")}}
    {{userReportString := cembed
        "title" "New Report"
        "description" (printf "<@%d> reported <@%d> in <#%d> for: `%s` \n Last 100 messages: <%s>" .User.ID $user.ID .Channel.ID $reason $s)
        "footer" (sdict "text" $reportGuide)}}
    {{dbSet 2000 "reportGuideBasic" $reportGuide}}
    {{dbSet 2000 (printf "userReport%d" .User.ID) $userReportString}}
    {{$x := sendMessageRetID $reports $userReportString}}
    {{addMessageReactions $reports $x "❌" "🛡️" "⚠️"}}
    User reported to the proper authorites!
    {{dbSet .User.ID "key" $secret}}
    {{sendDM (printf "User reported to the proper authorities! If you wish to cancel your report, simply type `-cancelr %d %s` in any channel.\n A reason is required." $x $secret)}}
{{end}}
