package models

var Data = []Service{
	{
		ID:          "1",
		Name:        "Data Processor",
		Description: "Processes data from various sources.",
		Versions: []Version{
			{ID: "1.1", Number: "1.0", CreatedBy: "Alice", CommitHash: "abc123"},
			{ID: "1.2", Number: "2.0", CreatedBy: "Bob", CommitHash: "def456"},
		},
	},
	{
		ID:          "2",
		Name:        "Resource Manager",
		Description: "Manages resource allocation and tracking.",
		Versions: []Version{
			{ID: "2.1", Number: "1.0", CreatedBy: "Charlie", CommitHash: "ghi789"},
		},
	},
	{
		ID:          "3",
		Name:        "User Authentication",
		Description: "Handles user authentication and security.",
		Versions: []Version{
			{ID: "3.1", Number: "1.0", CreatedBy: "Dave", CommitHash: "jkl012"},
			{ID: "3.2", Number: "1.1", CreatedBy: "Eve", CommitHash: "mno345"},
		},
	},
	{
		ID:          "4",
		Name:        "Payment Gateway",
		Description: "Processes payments and manages transaction security.",
		Versions: []Version{
			{ID: "4.1", Number: "1.0", CreatedBy: "Frank", CommitHash: "pqr678"},
			{ID: "4.2", Number: "2.0", CreatedBy: "Grace", CommitHash: "stu901"},
			{ID: "4.3", Number: "2.1", CreatedBy: "Hank", CommitHash: "vwx234"},
		},
	},
	{
		ID:          "5",
		Name:        "Email Service",
		Description: "Sends emails and handles email security.",
		Versions: []Version{
			{ID: "5.1", Number: "1.0", CreatedBy: "Ivy", CommitHash: "yz0123"},
		},
	},
	{
		ID:          "6",
		Name:        "File Storage",
		Description: "Stores files and handles file security.",
		Versions: []Version{
			{ID: "6.1", Number: "1.0", CreatedBy: "Jack", CommitHash: "123456"},
			{ID: "6.2", Number: "1.1", CreatedBy: "Kate", CommitHash: "789012"},
			{ID: "6.3", Number: "1.2", CreatedBy: "Liam", CommitHash: "123456"},
			{ID: "6.4", Number: "1.3", CreatedBy: "Mia", CommitHash: "789012"},
			{ID: "6.5", Number: "1.4", CreatedBy: "Noah", CommitHash: "123456"},
		},
	},
	{
		ID:          "7",
		Name:        "Notification Service",
		Description: "Sends notifications and handles notification security.",
		Versions: []Version{
			{ID: "7.1", Number: "1.0", CreatedBy: "Kate", CommitHash: "789012"},
			{ID: "7.2", Number: "1.1", CreatedBy: "Liam", CommitHash: "123456"},
		},
	},
	{
		ID:          "8",
		Name:        "API Gateway",
		Description: "Sends API requests and handles API security.",
		Versions: []Version{
			{ID: "8.1", Number: "1.0", CreatedBy: "Liam", CommitHash: "123456"},
			{ID: "8.2", Number: "1.1", CreatedBy: "Olivia", CommitHash: "789012"},
			{ID: "8.3", Number: "1.2", CreatedBy: "Emma", CommitHash: "123456"},
		},
	},
	{
		ID:          "9",
		Name:        "Caching Service",
		Description: "Caches data and handles caching security.",
		Versions: []Version{
			{ID: "9.1", Number: "1.0", CreatedBy: "Mia", CommitHash: "789012"},
		},
	},
	{
		ID:          "10",
		Name:        "Data Store",
		Description: "Stores data and handles data security.",
		Versions: []Version{
			{ID: "10.1", Number: "1.0", CreatedBy: "Natalie", CommitHash: "123456"},
			{ID: "10.2", Number: "1.1", CreatedBy: "Olivia", CommitHash: "789012"},
			{ID: "10.3", Number: "1.2", CreatedBy: "Penelope", CommitHash: "123456"},
			{ID: "10.4", Number: "1.3", CreatedBy: "Quinn", CommitHash: "789012"},
		},
	},
	{
		ID:          "11",
		Name:        "Data Analyzer",
		Description: "Analyzes data and handles data security.",
		Versions: []Version{
			{ID: "11.1", Number: "1.0", CreatedBy: "Olivia", CommitHash: "789012"},
			{ID: "11.2", Number: "1.1", CreatedBy: "Penelope", CommitHash: "123456"},
			{ID: "11.3", Number: "1.2", CreatedBy: "Quinn", CommitHash: "789012"},
		},
	},
	{
		ID:          "12",
		Name:        "Data Transformer",
		Description: "Transforms data and handles data security.",
		Versions: []Version{
			{ID: "12.1", Number: "1.0", CreatedBy: "Penelope", CommitHash: "123456"},
		},
	},
	{
		ID:          "13",
		Name:        "Data Aggregator",
		Description: "Aggregates data and handles data security.",
		Versions: []Version{
			{ID: "13.1", Number: "1.0", CreatedBy: "Quinn", CommitHash: "789012"},
			{ID: "13.2", Number: "1.1", CreatedBy: "Riley", CommitHash: "123456"},
			{ID: "13.3", Number: "1.2", CreatedBy: "Samantha", CommitHash: "789012"},
		},
	},
	{
		ID:          "14",
		Name:        "Data Enricher",
		Description: "Enriches data and handles data security.",
		Versions: []Version{
			{ID: "14.1", Number: "1.0", CreatedBy: "Riley", CommitHash: "123456"},
			{ID: "14.2", Number: "1.1", CreatedBy: "Samantha", CommitHash: "789012"},
		},
	},
	{
		ID:          "15",
		Name:        "Data Validator",
		Description: "Validates data and handles data security.",
		Versions: []Version{
			{ID: "15.1", Number: "1.0", CreatedBy: "Samantha", CommitHash: "789012"},
		},
	},
	{
		ID:          "16",
		Name:        "Data Formatter",
		Description: "Formats data and handles data security.",
		Versions: []Version{
			{ID: "16.1", Number: "1.0", CreatedBy: "Tegan", CommitHash: "123456"},
			{ID: "16.2", Number: "1.1", CreatedBy: "Uma", CommitHash: "789012"},
			{ID: "16.3", Number: "1.2", CreatedBy: "Violet", CommitHash: "123456"},
		},
	},
	{
		ID:          "17",
		Name:        "Data Extractor",
		Description: "Extracts data and handles data security.",
		Versions: []Version{
			{ID: "17.1", Number: "1.0", CreatedBy: "Uma", CommitHash: "789012"},
		},
	},
	{
		ID:          "18",
		Name:        "Data Loader",
		Description: "Loads data and handles data security.",
		Versions: []Version{
			{ID: "18.1", Number: "1.0", CreatedBy: "Violet", CommitHash: "123456"},
		},
	},
	{
		ID:          "19",
		Name:        "Data Uploader",
		Description: "Uploads data and handles data security.",
		Versions: []Version{
			{ID: "19.1", Number: "1.0", CreatedBy: "Wendy", CommitHash: "789012"},
		},
	},
	{
		ID:          "20",
		Name:        "Data Transformer",
		Description: "Transforms data and handles data security.",
		Versions: []Version{
			{ID: "20.1", Number: "1.0", CreatedBy: "Xander", CommitHash: "123456"},
			{ID: "20.2", Number: "1.1", CreatedBy: "Yara", CommitHash: "789012"},
			{ID: "20.3", Number: "1.2", CreatedBy: "Zara", CommitHash: "123456"},
			{ID: "20.4", Number: "1.3", CreatedBy: "Abby", CommitHash: "789012"},
		},
	},
	{
		ID:          "21",
		Name:        "Data Processor",
		Description: "Processes data and handles data security.",
		Versions: []Version{
			{ID: "21.1", Number: "1.0", CreatedBy: "Yara", CommitHash: "789012"},
		},
	},
	{
		ID:          "22",
		Name:        "Data Aggregator",
		Description: "Aggregates data and handles data security.",
		Versions: []Version{
			{ID: "22.1", Number: "1.0", CreatedBy: "Zara", CommitHash: "123456"},
		},
	},
	{
		ID:          "23",
		Name:        "Data Enricher",
		Description: "Enriches data and handles data security.",
		Versions: []Version{
			{ID: "23.1", Number: "1.0", CreatedBy: "Abby", CommitHash: "789012"},
			{ID: "23.2", Number: "1.1", CreatedBy: "Cara", CommitHash: "123456"},
			{ID: "23.3", Number: "1.2", CreatedBy: "Dara", CommitHash: "789012"},
			{ID: "23.4", Number: "1.3", CreatedBy: "Eva", CommitHash: "123456"},
			{ID: "23.5", Number: "1.4", CreatedBy: "Fara", CommitHash: "789012"},
			{ID: "23.6", Number: "1.5", CreatedBy: "Gara", CommitHash: "123456"},
		},
	},
	{
		ID:          "24",
		Name:        "Data Validator",
		Description: "Validates data and handles data security.",
		Versions: []Version{
			{ID: "24.1", Number: "1.0", CreatedBy: "Cara", CommitHash: "123456"},
		},
	},
	{
		ID:          "25",
		Name:        "Data Formatter",
		Description: "Formats data and handles data security.",
		Versions: []Version{
			{ID: "25.1", Number: "1.0", CreatedBy: "Dara", CommitHash: "789012"},
			{ID: "25.2", Number: "1.1", CreatedBy: "Eva", CommitHash: "123456"},
			{ID: "25.3", Number: "1.2", CreatedBy: "Fara", CommitHash: "789012"},
		},
	},
	{
		ID:          "26",
		Name:        "Data Extractor",
		Description: "Extracts data and handles data security.",
		Versions: []Version{
			{ID: "26.1", Number: "1.0", CreatedBy: "Eva", CommitHash: "123456"},
		},
	},
	{
		ID:          "27",
		Name:        "Data Loader",
		Description: "Loads data and handles data security.",
		Versions: []Version{
			{ID: "27.1", Number: "1.0", CreatedBy: "Fara", CommitHash: "789012"},
			{ID: "27.2", Number: "1.1", CreatedBy: "Gara", CommitHash: "123456"},
		},
	},
	{
		ID:          "28",
		Name:        "Data Uploader",
		Description: "Uploads data and handles data security.",
		Versions: []Version{
			{ID: "28.1", Number: "1.0", CreatedBy: "Gara", CommitHash: "123456"},
			{ID: "28.2", Number: "1.1", CreatedBy: "Hara", CommitHash: "789012"},
			{ID: "28.3", Number: "1.2", CreatedBy: "Iara", CommitHash: "123456"},
		},
	},
	{
		ID:          "29",
		Name:        "Data Transformer",
		Description: "Transforms data and handles data security.",
		Versions: []Version{
			{ID: "29.1", Number: "1.0", CreatedBy: "Hara", CommitHash: "789012"},
		},
	},
	{
		ID:          "30",
		Name:        "Data Processor",
		Description: "Processes data and handles data security.",
		Versions: []Version{
			{ID: "30.1", Number: "1.0", CreatedBy: "Iara", CommitHash: "123456"},
			{ID: "30.2", Number: "1.1", CreatedBy: "Jara", CommitHash: "789012"},
			{ID: "30.3", Number: "1.2", CreatedBy: "Kara", CommitHash: "123456"},
			{ID: "30.4", Number: "1.3", CreatedBy: "Lara", CommitHash: "789012"},
		},
	},
	{
		ID:          "31",
		Name:        "Data Aggregator",
		Description: "Aggregates data and handles data security.",
		Versions: []Version{
			{ID: "31.1", Number: "1.0", CreatedBy: "Jara", CommitHash: "789012"},
		},
	},
	{
		ID:          "32",
		Name:        "Data Enricher",
		Description: "Enriches data and handles data security.",
		Versions: []Version{
			{ID: "32.1", Number: "1.0", CreatedBy: "Kara", CommitHash: "123456"},
			{ID: "32.2", Number: "1.1", CreatedBy: "Lara", CommitHash: "789012"},
			{ID: "32.3", Number: "1.2", CreatedBy: "Mara", CommitHash: "123456"},
			{ID: "32.4", Number: "1.3", CreatedBy: "Nara", CommitHash: "789012"},
		},
	},
	{
		ID:          "33",
		Name:        "Data Validator",
		Description: "Validates data and handles data security.",
		Versions: []Version{
			{ID: "33.1", Number: "1.0", CreatedBy: "Lara", CommitHash: "789012"},
		},
	},
	{
		ID:          "34",
		Name:        "Data Formatter",
		Description: "Formats data and handles data security.",
		Versions: []Version{
			{ID: "34.1", Number: "1.0", CreatedBy: "Mara", CommitHash: "123456"},
			{ID: "34.2", Number: "1.1", CreatedBy: "Nara", CommitHash: "789012"},
			{ID: "34.3", Number: "1.2", CreatedBy: "Oara", CommitHash: "123456"},
			{ID: "34.4", Number: "1.3", CreatedBy: "Para", CommitHash: "789012"},
		},
	},
	{
		ID:          "35",
		Name:        "Data Extractor",
		Description: "Extracts data and handles data security.",
		Versions: []Version{
			{ID: "35.1", Number: "1.0", CreatedBy: "Nara", CommitHash: "789012"},
		},
	},
}
