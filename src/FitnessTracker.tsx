import React, { useState, useEffect } from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';
import DatePicker from 'react-datepicker';
import "react-datepicker/dist/react-datepicker.css";

interface Workout {
  date: Date;
  duration: number;
  intensity: number;
}

const FitnessTracker: React.FC = () => {
  const [workoutLog, setWorkoutLog] = useState<Workout[]>([]);
  const [displayedWorkouts, setDisplayedWorkouts] = useState<Workout[]>([]);
  const [newWorkout, setNewWorkout] = useState<Workout>({ date: new Date(), duration: 0, intensity: 0 });
  const [goalInMinutes, setGoalInMinutes] = useState<number>(0);
  const [startDateFilter, setStartDateFilter] = useState<Date | null>(new Date());
  const [endDateFilter, setEndDateFilter] = useState<Date | null>(new Date());

  useEffect(() => {
    filterWorkoutsByDate();
  }, [workoutLog, startDateFilter, endDateFilter]);

  const filterWorkoutsByDate = () => {
    if (!startDateFilter || !endDateFilter) {
      setDisplayedWorkouts(workoutLog);
      return;
    }
    const filtered = workoutLog.filter(workout => {
      const workoutDate = new Date(workout.date).getTime();
      return workoutDate >= new Date(startDateFilter).getTime() && workoutDate <= new Date(endDateFilter).getTime();
    });
    setDisplayedWorkouts(filtered);
  };

  const handleInputChange = (
    event: React.ChangeEvent<HTMLInputElement> | React.ChangeEvent<HTMLSelectElement>,
    fieldName: string
  ) => {
    setNewWorkout({
      ...newWorkworkout,
      [fieldName]: fieldName === "date" ? new Date(event.target.value) : parseInt(event.target.value)
    });
  };

  const addWorkoutToLog = () => {
    if (newWorkout.duration > 0 && newWorkout.intensity > 0) {
      setWorkoutLog([...workoutLog, newWorkout]);
      setNewWorkout({ date: new Date(), duration: 0, intensity: 0 });
    }
  };

  const achievedDuration = displayedWorkouts.reduce((total, { duration }) => total + duration, 0);
  const hasMetGoal = achievedDuration >= goalInMinutes;

  return (
    <div>
      <h2>Fitness Tracker</h2>
      <div>
        <label>
          Select Date:
          <DatePicker selected={newWorkout.date} onChange={(date: Date) => setNewWorkout({ ...newWorkout, date })} />
        </label>
        <label>
          Duration (in minutes):
          <input type="number" value={newWorkout.duration} onChange={(event) => handleInputChange(event, 'duration')} />
        </label>
        <label>
          Intensity (1-10):
          <input type="number" min="1" max="10" value={newWorkout.intensity} onChange={(event) => handleInputChange(event, 'intensity')} />
        </label>
        <button onClick={addWorkoutToLog}>Log Workout</button>
      </div>
      <div>
        <label>
          Set your fitness goal (in minutes):
          <input type="number" value={goalInMinutes} onChange={(event) => setGoalInMinutes(parseInt(event.target.value))} />
        </label>
        {hasMetGoal && <div>Congratulations! You've achieved your goal!</div>}
      </div>
      <div>
        <label>
          Start Date:
          <DatePicker selected={startDateFilter} onChange={setStartDateFilter} />
        </label>
        <label>
          End Date:
          <DatePicker selected={endDateFilter} onChange={setEndDateFilter} />
        </label>
        <button onClick={filterWorkoutsByDate}>Filter Logs</button>
      </div>
      <LineChart width={500} height={300} data={displayedWorkouts}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="date" />
        <YAxis />
        <Tooltip />
        <Legend />
        <Line type="monotone" dataKey="duration" stroke="#8884d8" />
        <Line type="monotone" dataKey="intensity" stroke="#82ca9d" />
      </LineChart>
    </div>
  );
};

export default FitnessTracker;