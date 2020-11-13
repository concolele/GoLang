# GoLang
GoLang + Graphql + Postgre

■Import lib

go get github.com/graphql-go/graphql
go get github.com/graphql-go/handler
go get github.com/lib/pq


■Postgres database

Install Postgres
User: postgre
Pass: admin123
Create db LearnGoLang

■Create table, insert data

INSERT INTO skill(name, desc_skill) Values ('dev', '1')

INSERT INTO member(name, age, skill_id) Values ('vinh', '27', 1)
CREATE TABLE IF NOT EXISTS member
(
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    age int,
    skill_id int,
)

CREATE TABLE IF NOT EXISTS skill
(
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    desc_skill varchar(100) NOT NULL,
)

■Run Go

cmd + [go run LearnGoGraph.go]

Browser: http://localhost:8080/graphql

1. Get all Member
{
  members {
    id
    name
    age
    skill_id {
      id
    }
  }
}
-----------------------------------
2. Get member with id
{
  member(id: 2) {
    id
    name
    age
    skill_id {
      id
    }
  }
}
-----------------------------------
3. Create member
mutation {
  createMember(name: "fb", age: 1, skill_id: 1) {
    name
    age
    skill_id {
      id
    }
  }
}
-----------------------------------
4. Update member
mutation {
  updateMember(id: 1, name: "fb", age: 1, skill_id: 1) {
    id
    name
    age
    skill_id {
      id
    }
  }
}
-----------------------------------
5. Delete member
mutation {
  deleteMember(id: 1) {
    id
  }
}
-----------------------------------
6. Select all skill
{
  skills {
    id
    name
    desc_skill
  }
}
-----------------------------------
7. Select skill with id
{
  skill(id: 1) {
    id
    name
    desc_skill
  }
}
-----------------------------------
8. Create new skill
mutation {
  createSkill(name: "manager", desc_skill: "2") {
    name
    desc_skill
  }
}
-----------------------------------
9. Update skill
mutation {
  updateSkill(id: 1, name: "brse", desc_skill: "1") {
    id
    name
    desc_skill
  }
}
-----------------------------------
10. Delete skill
mutation {
  deleteSkill(id: 1) {
    id
  }
}

★Ref：
https://github.com/sohelamin/graphql-postgres-go/blob/master/README.md
















