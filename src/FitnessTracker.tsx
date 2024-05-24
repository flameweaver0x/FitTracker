import React, { useState, useEffect } from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';
import DatePicker from 'react-datepicker';
import "react-datepicker/dist/react-datepicker.css";

interface IWorkout {
  date: Date;
  duration: number;
  intensity: number;
}

const FitnessTracker: React.FC = () => {
  const [workoutLog, setWorkoutLog] = useState<IWorkout[]>([]);
  const [filteredWorkouts, setFilteredWorkouts] = useState<IWorkout[]>([]);
  const [workoutForm, setWorkoutForm] = useState<IWorkout>({ date: new Date(), duration: 0, intensity: 0 });
  const [goalMinutes, setGoalMinutes] = useState<number>(0);
  const [startDate, setStartDate] = useState<Date | null>(new Date());
  const [endDate, setEndDate] = useState<Date | null>(new Date());

  useEffect(() => {
    filterWorkoutsByDateRange();
  }, [workoutLog, startDate, endDate]);

  const filterWorkoutsByDateRange = () => {
    if (!startDate || !endDate) {
      setFilteredWorkouts(workoutLog);
      return;
    }
    const filtered = workoutLog.filter(workout => {
      const workoutDate = new Date(workout.date).getTime();
      return workoutDate >= new Date(startDate).getTime() && workoutDate <= new Date(endDate).getTime();
    });
    setFilteredWorkouts(filtered);
  };

  const handleWorkoutInputChange = (event: React.ChangeEvent<HTMLInputElement> | React.ChangeEvent<HTMLSelectElement>, fieldName: string) => {
    setWorkoutForm({
      ...workoutForm,
      [fieldName]: fieldName === "date" ? new Date(event.target.value) : parseInt(event.target.value)
    });
  };

  const addWorkoutToLog = () => {
    if (workoutForm.duration > 0 && workoutForm.intensity > 0) {
      setWorkoutLog([...workoutLog, workoutForm]);
      setWorkoutForm({ date: new Date(), duration: 0, intensity: 0 });
    }
  };

  const accumulatedDuration = filteredWorkouts.reduce((total, current) => total + current.duration, 0);

  const hasAchievedGoal = accumulatedDuration >= goalMinutes;

  return (
    <div>
      <h2>Fitness Tracker</h2>
      <div>
        <label>
          Select Date:
          <DatePicker selected={workoutForm.date} onChange={(date: Date) => setWorkoutForm({ ...workoutForm, date })} />
        </label>
        <label>
          Duration (in minutes):
          <input type="number" value={workoutForm.duration} onChange={(event) => handleWorkoutInputChange(event, 'duration')} />
        </label>
        <label>
          Intensity (1-10):
          <input type="number" min="1" max="10" value={workoutForm.intensity} onChange={(event) => handleWorkoutInputChange(event, 'intensity')} />
        </label>
        <button onClick={addWorkoutToLog}>Log Workout</button>
      </div>
      <div>
        <label>
          Set your fitness goal (in minutes):
          <input type="number" value={goalMinutes} onChange={(event) => setGoalMinutes(parseInt(event.target.value))} />
        </label>
        {hasAchievedGoal && <div>Congratulations! You've achieved your goal!</div>}
      </div>
      <div>
        <label>
          Start Date:
          <DatePicker selected={startDate} onChange={(date: Date) => setStartDate(date)} />
        </label>
        <label>
          End Date:
          <DatePicker selected={endDate} onChange={(date: Date) => setEndDate(date)} />
        </label>
        <button onClick={filterWorkoutsByDateRange}>Filter logs</button>
      </div>
      <LineChart width={500} height={300} data={filteredWorkouts}>
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