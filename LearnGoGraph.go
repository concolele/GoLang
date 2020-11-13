package main

import (
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
	"net/http"
)

type Member struct {
   ID           int        `json:"id"`
   Name         string     `json:"name"`
   Age          int        `json:"age"`
   Skill_ID     int        `json:"skill_id"`
}

type Skill struct {
   ID          int        `json:"id"`
   Name        string     `json:"name"`
   Desc_skill  string     `json:"desc_skill"`
}


// type Author struct {
// 	ID        int       `json:"id"`
// 	Name      string    `json:"name"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// type Post struct {
// 	ID        int       `json:"id"`
// 	Title     string    `json:"title"`
// 	Content   string    `json:"content"`
// 	AuthorID  int       `json:"author_id"`
// 	CreatedAt time.Time `json:"created_at"`
// }

const (
	user     = "postgres"
	password = "admin123"
	dbname   = "LearnGoLang"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	defer db.Close()

	skillType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Skill",
		Description: "The skill",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*Skill); ok {
						return skill.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*Skill); ok {
						return skill.Name, nil
					}

					return nil, nil
				},
			},
			"desc_skill": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The desc of the skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*Skill); ok {
						return skill.Desc_skill, nil
					}

					return nil, nil
				},
			},
		},
	})

	memberType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Member",
		Description: "A Member",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the member.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if member, ok := p.Source.(*Member); ok {
						return member.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the member.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if member, ok := p.Source.(*Member); ok {
						return member.Name, nil
					}

					return nil, nil
				},
			},
			"age": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The age of the member.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if member, ok := p.Source.(*Member); ok {
						return member.Age, nil
					}

					return nil, nil
				},
			},
			"skill_id": &graphql.Field{
				Type: skillType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if member, ok := p.Source.(*Member); ok {
						skill := &Skill{}
						err = db.QueryRow("select id, name, desc_skill from skill where id = $1", member.Skill_ID).Scan(&skill.ID, &skill.Name, &skill.Desc_skill)
						checkErr(err)

						return skill, nil
					}

					return nil, nil
				},
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"member": &graphql.Field{
				Type:        memberType,
				Description: "Get an member.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					member := &Member{}
					err = db.QueryRow("select id, name, age, skill_id from member where id = $1", id).Scan(&member.ID, &member.Name, &member.Age, &member.Skill_ID)
					checkErr(err)

					return member, nil
				},
			},
			"members": &graphql.Field{
				Type:        graphql.NewList(memberType),
				Description: "List of members.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT id, name, age, skill_id FROM member")
					checkErr(err)
					var members []*Member

					for rows.Next() {
						member := &Member{}

						err = rows.Scan(&member.ID, &member.Name, &member.Age, &member.Skill_ID)
						checkErr(err)
						members = append(members, member)
					}

					return members, nil
				},
			},
			"skill": &graphql.Field{
				Type:        skillType,
				Description: "Get a skill.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					skill := &Skill{}
					err = db.QueryRow("select id, name, desc_skill from skill where id = $1", id).Scan(&skill.ID, &skill.Name, &skill.Desc_skill)
					checkErr(err)

					return skill, nil
				},
			},
			"skills": &graphql.Field{
				Type:        graphql.NewList(skillType),
				Description: "List of skills.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT id, name, desc_skill FROM skill")
					checkErr(err)
					var skills []*Skill

					for rows.Next() {
						skill := &Skill{}

						err = rows.Scan(&skill.ID, &skill.Name, &skill.Desc_skill)
						checkErr(err)
						skills = append(skills, skill)
					}

					return skills, nil
				},
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// Member
			"createMember": &graphql.Field{
				Type:        memberType,
				Description: "Create new member",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"skill_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					age, _ := params.Args["age"].(int)
					skill_id, _ := params.Args["skill_id"].(int)

					var lastInsertId int
					err = db.QueryRow("INSERT INTO member(name, age, skill_id) VALUES($1, $2, $3) returning id;", name, age, skill_id).Scan(&lastInsertId)
					checkErr(err)

					newMember := &Member{
						ID:        lastInsertId,
						Name:      name,
						Age:       age,
						Skill_ID:   skill_id,
					}

					return newMember, nil
				},
			},
			"updateMember": &graphql.Field{
				Type:        memberType,
				Description: "Update an member",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"skill_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					age, _ := params.Args["age"].(int)
					skill_id, _ := params.Args["skill_id"].(int)

					stmt, err := db.Prepare("UPDATE member SET name = $1, age = $2, skill_id = $3 WHERE id = $4")
					checkErr(err)

					_, err2 := stmt.Exec(name, age, skill_id, id)
					checkErr(err2)

					newMember := &Member{
						ID:    id,
						Name:  name,
						Age: age,
						Skill_ID: skill_id,
					}

					return newMember, nil
				},
			},
			"deleteMember": &graphql.Field{
				Type:        memberType,
				Description: "Delete an member",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					stmt, err := db.Prepare("DELETE FROM member WHERE id = $1")
					checkErr(err)

					_, err2 := stmt.Exec(id)
					checkErr(err2)

					return nil, nil
				},
			},
			// Skill
			"createSkill": &graphql.Field{
				Type:        skillType,
				Description: "Create new skill",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"desc_skill": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					desc_skill, _ := params.Args["desc_skill"].(string)

					var lastInsertId int
					err = db.QueryRow("INSERT INTO skill(name, desc_skill) VALUES($1, $2) returning id;", name, desc_skill).Scan(&lastInsertId)
					checkErr(err)

					newSkill := &Skill{
						ID:              lastInsertId,
						Name:            name,
						Desc_skill:      desc_skill,
					}

					return newSkill, nil
				},
			},
			"updateSkill": &graphql.Field{
				Type:        skillType,
				Description: "Update a skill",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"desc_skill": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					desc_skill, _ := params.Args["desc_skill"].(string)

					stmt, err := db.Prepare("UPDATE skill SET name = $1, desc_skill = $2 WHERE id = $3")
					checkErr(err)

					_, err2 := stmt.Exec(name, desc_skill, id)
					checkErr(err2)

					newSkill := &Skill{
						ID:             id,
						Name:           name,
						Desc_skill:     desc_skill,
					}

					return newSkill, nil
				},
			},
			"deleteSkill": &graphql.Field{
				Type:        skillType,
				Description: "Delete a skill",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					stmt, err := db.Prepare("DELETE FROM skill WHERE id = $1")
					checkErr(err)

					_, err2 := stmt.Exec(id)
					checkErr(err2)

					return nil, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}