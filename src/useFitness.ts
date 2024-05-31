import { useState, useEffect } from 'react';
import axios from 'axios';

interface Workout {
  id: string;
  date: string;
  exercises: Exercise[];
}

interface Exercise {
  name: string;
  sets: number;
  reps: number;
  weight: number;
}

const useWorkoutData = () => {
  const [workoutLogs, setWorkoutLogs] = useState<Workout[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [fetchError, setFetchError] = useState<string | null>(null);

  const apiBaseUrl = process.env.REACT_APP_BACKEND_URL;

  const handleLoadingState = (loading: boolean) => {
    setIsLoading(loading);
  };

  const handleFetchError = (message: string) => {
    setFetchError(message);
    handleLoadingState(false);
  };

  const updateWorkoutLogs = (data: Workout[]) => {
    setWorkoutLogs(data);
    handleLoadingState(false);
  };

  const fetchWorkoutLogs = async (): Promise<void> => {
    handleLoadingState(true);
    try {
      const { data } = await axios.get(`${apiBaseUrl}/workouts`);
      updateWorkoutLogs(data);
    } catch (error) {
      handleFetchError('Failed to fetch workout logs');
    }
  };

  const addExerciseToWorkout = async (workoutId: string, newExercise: Exercise): Promise<void> => {
    try {
      await axios.post(`${apiBaseUrl}/workouts/${workoutId}/exercises`, newExercise);
      fetchWorkingLogsWrapper();
    } catch (error) {
      handleFetchError('Failed to add exercise to workout log');
    }
  };

  const fetchWorkingLogsWrapper = async () => {
    await fetchWorkoutLogs(); // Optionally, you can handle UI updates here if needed
  };

  const fetchRecommendedWorkouts = async (): Promise<Workout[] | undefined> => {
    try {
      const { data } = await axios.get(`${apiBaseUrl}/recommendations`);
      return data;
    } catch (error) {
      handleFetchError('Failed to fetch workout recommendations');
      return undefined;
    }
  };

  useEffect(() => {
    fetchWorkingLogsWrapper();
  }, []);

  return {
    workoutLogs,
    fetchWorkoutLogs,
    addExerciseToWorkout,
    fetchRecommendedWorkouts,
    isLoading,
    fetchError,
  };
};

export default useWorkoutData;