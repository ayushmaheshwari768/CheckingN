package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/middleware"

	"github.com/gorilla/mux"
)

func (h *Handlers) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment middleware.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		log.Println("ERROR [handlers/appointmentHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.db.AddAppointment(appointment); err != nil {
		log.Printf("Couldn't Create Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}


func (h *Handlers) EditAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := mux.Vars(r)["id"]
	var appointment middleware.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		log.Println("ERROR [handlers/appointmentHandlers.go] Couldn't get data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.db.UpdateAppointment(appointmentId, appointment); err != nil {
		log.Printf("Couldn't Edit Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusConflict)
		return
	}

	log.Println("Successfully edit appointment")

}

func (h *Handlers) CancelAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := mux.Vars(r)["id"]

	if err := h.db.DeleteAppointment(appointmentId); err != nil {
		log.Printf("Couldn't Delete Appointment Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) ViewAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := mux.Vars(r)["id"]

	appointment, err := h.db.GetAppointment(appointmentId)
	if err != nil {
		log.Printf("Couldn't View Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, appointment.TutorID)
	fmt.Fprintln(w, appointment.StudentID)
	fmt.Fprintln(w, appointment.CourseCode)
	fmt.Fprintln(w, appointment.MeetingLocation)
	fmt.Fprintln(w, appointment.StartTime.String())
	fmt.Fprintln(w, appointment.EndTime.String())
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) ViewAllTutorAppointment(w http.ResponseWriter, r *http.Request) {
	tutorId := mux.Vars(r)["id"]

	appointments, err := h.db.GetAppointmentsForTutor(tutorId)
	if err != nil {
		log.Printf("Couldn't Get Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, appointment := range appointments {
		fmt.Fprintln(w, appointment.TutorID)
		fmt.Fprintln(w, appointment.StudentID)
		fmt.Fprintln(w, appointment.CourseCode)
		fmt.Fprintln(w, appointment.MeetingLocation)
		fmt.Fprintln(w, appointment.StartTime.String())
		fmt.Fprintln(w, appointment.EndTime.String())
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) ViewAllStudentAppointment(w http.ResponseWriter, r *http.Request) {
	studentId := mux.Vars(r)["id"]

	appointments, err := h.db.GetAppointmentsForStudent(studentId)
	if err != nil {
		log.Printf("Couldn't Get Appointment: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, appointment := range appointments {
		fmt.Fprintln(w, appointment.TutorID)
		fmt.Fprintln(w, appointment.StudentID)
		fmt.Fprintln(w, appointment.CourseCode)
		fmt.Fprintln(w, appointment.MeetingLocation)
		fmt.Fprintln(w, appointment.StartTime.String())
		fmt.Fprintln(w, appointment.EndTime.String())
	}
	w.WriteHeader(http.StatusOK)
}

/*Test case

For Make Appointment, Post Method
{
"tutor_id": "616f427a3b3c421b64576b51",
"student_id": "6171c45e712f8abc5340a8e8",
"course_code": "CS146",
"meeting_location": "Zoom",
"start_time": "2021-10-25T09:00:00+00:00",
"end_time": "2021-10-25T10:00:00+00:00"
}

For Delete Appointment, Delete Method
http://localhost:8080/api/appointment/{id}
replace {id} with Appointment _id


For Edit Appointment, Patch Method
http://localhost:8080/api/appointment/{id}
replace {id} with Appointment _id

{
"tutor_id": "616f427a3b3c421b64576b51",
"student_id": "6171c45e712f8abc5340a8e8",
"course_code": "CS146",
"meeting_location": "Zoom",
"start_time": "2021-10-27T14:00:00+00:00",
"end_time": "2021-10-27T15:00:00+00:00"
}

For View Appointment, Get Method
http://localhost:8080/api/appointment/{id}
replace {id} with Appointment _id

For View All Appointment, Get Method
http://localhost:8080/api/appointment/tutor{id}
http://localhost:8080/api/appointment/student/{id}
replace {id} with Student or Tutor _id
*/
