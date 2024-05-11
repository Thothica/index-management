username=""
password=""
endpoint=""
pipeline="-pipeline"
curl -XPUT "https://$endpoint:9200/_ingest/pipeline/$pipeline" \
	-u "$username:$password" \
	-k \
	-H "Content-Type: application/json" \
	-d'
	{
		"description": "embedding pipeline",
		"processors": [
			{
			"text_embedding": {
				"model_id": "AbDZGo8BB3UUeZ_94CHA",
				"field_map": {
					"Raw_Response": "Raw_Response_embedding"
					}
				}
			}
		]
	}'
