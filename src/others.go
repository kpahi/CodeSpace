//
// Author: Vinhthuy Phan, 2018
//
package main

import (
	"fmt"
	"net/http"
)

//-----------------------------------------------------------------------------------
func teacher_gets_passcodeHandler(w http.ResponseWriter, r *http.Request, who string, uid int) {
	fmt.Fprintf(w, Passcode)
}

//-----------------------------------------------------------------------------------
func testHandler(w http.ResponseWriter, r *http.Request, who string, uid int) {
	// Show content of boards
	fmt.Println("Students:", len(Students))
	for uid, st := range Students {
		fmt.Printf("Uid: %d has %d pages. Status: %d\n", uid, len(st.Boards), st.SubmissionStatus)
		for i := 0; i < len(st.Boards); i++ {
			b := st.Boards[i]
			fmt.Printf("Attempts: %d, Filename: %s, Pid: %d, Answer: %s, len of content: %d\n",
				b.Attempts, b.Filename, b.Pid, b.Answer, len(b.Content))
		}
	}

	fmt.Printf("WorkingSubs: %d entries", len(WorkingSubs))
	for i := 0; i < len(WorkingSubs); i++ {
		fmt.Println(WorkingSubs[i].Sid, WorkingSubs[i].Uid, WorkingSubs[i].Pid, WorkingSubs[i].Priority)
		fmt.Println(WorkingSubs[i].Content)
	}
	fmt.Println()

	fmt.Println("ActiveProblems:", ActiveProblems)
	for pid, v := range ActiveProblems {
		fmt.Println(pid, v.Active, "Answers:", v.Answers, "Attempts:", v.Attempts)
		fmt.Println(pid, v.Info.Filename, v.Info.Merit, v.Info.Effort, v.Info.Attempts, v.Info.NextIfCorrect, v.Info.NextIfIncorrect)
	}
	fmt.Fprintf(w, Passcode)
}

//-----------------------------------------------------------------------------------