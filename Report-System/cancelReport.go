{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command manages and takes care of the cancellation requests.
    
    Recommended Trigger type and trigger: Regex; \A-c(ancel)?r(eport)?(\s+|\z)

    Credit: ye olde boi#7325 U-ID:665243449405997066
    Contributors@ Devonte#0745 U-ID:622146791659405313
*/}}


{{/*CONFIG AREA START*/}}

{{$reports := 750730537571975298}} {{/*The channel where your reports are logged into.*/}}
{{$reportDiscussion := 750099460314628176}} {{/*Your channel where users talk to staff*/}}

{{/*CONFIG AREA END*/}}


{{/*ACTUAL CODE DO NOT TOUCH UNLESS YOU KNOW WHAT YOU ARE DOING*/}}
{{if not (ge (len .CmdArgs) 3)}}
    ```{{.Cmd}} <Message:ID> <Key:Text> <Reason:Text>```
    Not enough arguments passed.
{{else}}
    {{$dbValue := (dbGet .User.ID "key").Value|str}}
    {{$reportMessage := ((index .CmdArgs 0)|toInt64)}}
    {{$reportMessageContent := (getMessage $reports $reportMessage).Content}}
    {{if (reFind `\A<@!?\d{17,19}>` $reportMessageContent)}} 
            {{if eq "used" $dbValue}}
                Your latest report has already been cancelled!
            {{else}}
            {{if eq (index .CmdArgs 1|str) $dbValue}}
                {{if ge (len .CmdArgs) 3}}
                    {{$reason := joinStr " " (slice .CmdArgs 2)}}
                    {{$userReportString := (dbGet 2000 (printf "UserCancel%d" .User.ID)).Value}}
                    {{$cancelGuide := (printf "\nDeny request with 🚫, accept with ✅, or request more information with ⚠️")}}
                    {{dbSet 2000 "cancelGuideBasic" $cancelGuide}}
                    {{$userCancelString := cembed
                        "title" "Cancel Report Request"
                        "description" (print "<@%d> requested cancellation of this report\n**Reason**\n`%s`" .User.ID $reason)
                        "footer" (sdict "text" $cancelGuide)}}
                    {{dbSet 2000 (printf "userCancel%d" .User.ID) $userCancelString}}
                    {{editMessage $reports $reportMessage (complexMessageEdit "embed" $userCancelString)}}
                    Cancellation requested.
                    {{deleteAllMessageReactions $reports $reportMessage}}
                    {{addMessageReactions $reports $reportMessage "🚫" "✅" "⚠️"}}
                    {{dbSet .User.ID "key" "used"}}
                {{end}}
            {{else}}
                Invalid key provided!
            {{end}}
        {{end}}
        {{else}}
            You are not the author of this report!
    {{end}}
{{end}}
