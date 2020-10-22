{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command is basically the native report-command, but adds some back-end functionalites in order for the rest to work :)


    Recommended Trigger type and trigger: Regex; \A-r(eport)?u(ser)?(\s+|\z)

    Credit: ye olde boi#7325 U-ID:665243449405997066
*/}}
{{/*CONFIG AREA START*/}}
{{$reports := 750730537571975298}} {{/*The channel where your reports are logged into.*/}}
{{/*CONFIG AREA END*/}}
{{if not (ge (len .CmdArgs) 2)}}
    ```{{.Cmd}} <User:Mention/ID> <Reason:Text>```
    Not enough arguments passed.
{{else}}
    {{$secret := adjective}}
    {{$s := execAdmin "log"}}
    {{$user := userArg (index .CmdArgs 0)}}
    {{$reason := joinStr " " (slice .CmdArgs 1)}}
    {{$reportGuide := (printf "\nDismiss report with ❌, take action with 🛡️, or request more background information with ⚠️")}}
    {{$userReportString := (printf  "<@%d> reported <@%d> in <#%d> for: `%s` \n Last 100 messages: <%s>" .User.ID $user.ID .Channel.ID $reason $s)}}
    {{dbSet 2000 "reportGuideBasic" $reportGuide}}
    {{dbSet 2000 (print "userReport-" .User.ID) $userReportString}}
    {{$x := sendMessageRetID $reports (printf "%s %s" $userReportString $reportGuide)}}
    {{addMessageReactions $reports $x "❌" "🛡️" "⚠️"}}
    User reported to the proper authorites!
    {{dbSet .User.ID "key" $secret}}
    {{sendDM (printf "User reported to the proper authorities! If you wish to cancel your report, simply type `-cancelr %d %s` in any channel.\n A reason is required." $x $secret)}}
{{end}}