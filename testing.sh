timestamp="$(date)"
head_commit_msg="abcd"
current_tag="v1.0"

req_body=$(cat <<EOF
{
    "content": "",
    "embeds": [
    {
        "title": "Aztemarket $current_tag was released!",
        "description": "$head_commit_msg",
        "color": 16711680,
        "fields": [
        {
            "name": "At",
            "value": "$timestamp"
        }
        ]
    }
    ]
} 
EOF
)

echo $req_body

curl -H "Content-Type: application/json" -d "$req_body" https://discord.com/api/webhooks/1246564340006518816/PkrvCxW96rFJQyLMJOQOfLTdOrrZLK_HmFOtrg8dGlwEtVlow4jrUZ9WM8pxU76Cr13i