package posts

type Post struct {
	FileName    string `json:"image_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorEmail string `json:"author_email"`
	AuthorName  string `json:"author_name"`
	AuthorRole  string `json:"author_role"`
}

var (
	AllPosts = []Post{
		{
			FileName:    "0.webp",
			Title:       "How to Build a Skeleton UI",
			Description: "A comprehensive guide to creating a Skeleton UI with modern tools.",
			AuthorEmail: "jane.doe@example.com",
			AuthorName:  "Jane Doe",
			AuthorRole:  "Frontend Developer",
		},
		{
			FileName:    "1.webp",
			Title:       "Mastering Go for Backend Dev",
			Description: "Learn advanced Go techniques for building high-performance backends.",
			AuthorEmail: "john.smith@example.com",
			AuthorName:  "John Smith",
			AuthorRole:  "Backend Engineer",
		},
		{
			FileName:    "2.webp",
			Title:       "Dynamic UIs with HTMX",
			Description: "Explore the power of HTMX for creating dynamic and reactive UIs.",
			AuthorEmail: "sarah.lee@example.com",
			AuthorName:  "Sarah Lee",
			AuthorRole:  "Fullstack Developer",
		},
		{
			FileName:    "3.webp",
			Title:       "Tailwind CSS: Best Practices",
			Description: "Optimize your development workflow with Tailwind CSS tips and tricks.",
			AuthorEmail: "emma.watson@example.com",
			AuthorName:  "Emma Watson",
			AuthorRole:  "UI/UX Designer",
		},
		{
			FileName:    "4.webp",
			Title:       "Building Accessible Web Apps",
			Description: "A practical approach to making your web applications accessible.",
			AuthorEmail: "liam.jones@example.com",
			AuthorName:  "Liam Jones",
			AuthorRole:  "Accessibility Advocate",
		},
		{
			FileName:    "5.webp",
			Title:       "Introduction to Leapkit",
			Description: "Get started with Leapkit and supercharge your HTMX applications.",
			AuthorEmail: "mia.taylor@example.com",
			AuthorName:  "Mia Taylor",
			AuthorRole:  "Software Engineer",
		},
		{
			FileName:    "6.webp",
			Title:       "Scaling Applications with Go",
			Description: "Best practices for building scalable applications in Go.",
			AuthorEmail: "noah.williams@example.com",
			AuthorName:  "Noah Williams",
			AuthorRole:  "DevOps Specialist",
		},
		{
			FileName:    "7.webp",
			Title:       "HTMX vs Traditional JS",
			Description: "A comparison of HTMX and traditional frameworks like React and Angular.",
			AuthorEmail: "ava.brown@example.com",
			AuthorName:  "Ava Brown",
			AuthorRole:  "Frontend Architect",
		},
		{
			FileName:    "8.webp",
			Title:       "Real-Time Data in Web Applications",
			Description: "Techniques for implementing real-time data updates in web apps.",
			AuthorEmail: "oliver.johnson@example.com",
			AuthorName:  "Oliver Johnson",
			AuthorRole:  "Realtime Data Engineer",
		},
		{
			FileName:    "9.webp",
			Title:       "Performance Tuning for Web Apps",
			Description: "Tips and tricks for improving the performance of your web applications.",
			AuthorEmail: "sophia.martin@example.com",
			AuthorName:  "Sophia Martin",
			AuthorRole:  "Performance Engineer",
		},
		{
			FileName:    "10.webp",
			Title:       "The Art of Responsive Design",
			Description: "Learn how to create responsive web designs that adapt seamlessly to any device.",
			AuthorEmail: "alex.smith@example.com",
			AuthorName:  "Alex Smith",
			AuthorRole:  "UI/UX Designer",
		},
		{
			FileName:    "11.webp",
			Title:       "Optimizing Web Performance",
			Description: "Techniques to enhance your web application's performance and speed.",
			AuthorEmail: "linda.jones@example.com",
			AuthorName:  "Linda Jones",
			AuthorRole:  "Web Performance Engineer",
		},
	}
)
