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
  const [workoutHistory, setWorkoutHistory] = useState<Workout[]>([]);
  const [visibleWorkouts, setVisibleWorkouts] = useState<Workout[]>([]);
  const [currentWorkoutEntry, setCurrentWorkoutEntry] = useState<Workout>({ date: new Date(), duration: 0, intensity: 0 });
  const [fitnessGoalMinutes, setFitnessGoalMinutes] = useState<number>(0);
  const [filterStartDate, setFilterStartDate] = useState<Date | null>(new Date());
  const [filterEndDate, setFilterEndDate] = useState<Date | null>(new Date());

  useEffect(() => {
    applyDateRangeFilter();
  }, [workoutHistory, filterStartDate, filterEndDate]);

  const applyDateRangeFilter = () => {
    if (!filterStartDate || !filterEndDate) {
      setVisibleWorkouts(workoutHistory);
      return;
    }
    const filteredWorkouts = workoutHistory.filter(workout => {
      const workoutDate = new Date(workout.date).getTime();
      return workoutDate >= new Date(filterStartDate).getTime() && workoutDate <= new Date(filterEndDate).getTime();
    });
    setVisibleWorkouts(filteredWorkouts);
  };

  const handleWorkoutChange = (event: React.ChangeEvent<HTMLInputElement> | React.ChangeEvent<HTMLSelectElement>, fieldName: string) => {
    setCurrentWorkoutEntry({
      ...currentWorkoutEntry,
      [fieldName]: fieldName === "date" ? new Date(event.target.value) : parseInt(event.target.value)
    });
  };

  const logNewWorkout = () => {
    if (currentWorkoutEntry.duration > 0 && currentWorkoutEntry.intensity > 0) {
      setWorkoutHistory([...workoutHistory, currentWorkoutEntry]);
      setCurrentWorkoutEntry({ date: new Date(), duration: 0, intensity: 0 });
    }
  };

  const totalDurationAchieved = visibleWorkouts.reduce((total, workout) => total + workout.duration, 0);

  const isGoalAchieved = totalDurationAchieved >= fitnessGoalMinutes;

  return (
    <div>
      <h2>Fitness Tracker</h2>
      <div>
        <label>
          Select Date:
          <DatePicker selected={currentWorkoutEntry.date} onChange={(date: Date) => setCurrentWorkoutEntry({ ...currentWorkoutEntry, date })} />
        </label>
        <label>
          Duration (in minutes):
          <input type="number" value={currentWorkoutEntry.duration} onChange={(event) => handleWorkoutChange(event, 'duration')} />
        </label>
        <label>
          Intensity (1-10):
          <input type="number" min="1" max="10" value={currentWorkoutEntry.intensity} onChange={(event) => handleWorkoutChange(event, 'intensity')} />
        </label>
        <button onClick={logNewWorkout}>Log Workout</button>
      </div>
      <div>
        <label>
          Set your fitness goal (in minutes):
          <input type="number" value={fitnessGoalMinutes} onChange={(event) => setFitnessGoalMinutes(parseInt(event.target.value))} />
        </label>
        {isGoalAchieved && <div>Congratulations! You've achieved your goal!</div>}
      </div>
      <div>
        <label>
          Start Date:
          <DatePicker selected={filterStartDate} onChange={(date: Date) => setFilterStartDate(date)} />
        </label>
        <label>
          End Date:
          <DatePicker selected={filterEndDate} onChange={(date: Date) => setFilterEndDate(date)} />
        </label>
        <button onClick={applyDateRangeFilter}>Filter Logs</button>
      </div>
      <LineChart width={500} height={300} data={visibleWorkouts}>
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