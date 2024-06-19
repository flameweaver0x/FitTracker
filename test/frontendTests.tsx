import React, { useState, useEffect, useCallback } from 'react';
import debounce from 'lodash/debounce';

const FitnessTracker: React.FC = () => {
  const [steps, setSteps] = useState(0);
  const [realTimeUpdates, setRealTimeUpdates] = useState('');

  // Assuming fetchSteps could evolve to need argments for caching, implementing a basic cache system
  const cache: { [key: string]: any } = {};

  const memoizedFetchSteps = useCallback(async () => {
    const cacheKey = 'fetchSteps';
    if (cache[cacheKey]) {
      setRealTimeUpdates(`Updated Steps: ${cache[cacheKey].steps}`);
      return;
    }
    const response = await fetch(`${process.env.REACT_APP_API_URL}/steps`);
    const data = await response.json();
    cache[cacheKey] = data; // Store response in cache
    setRealTimeUpdates(`Updated Steps: ${data.steps}`);
  }, [setRealTimeUpdates]);

  // Debounced function to fetch steps, waiting until user has stopped triggering fetch for 500ms
  const debouncedFetchSteps = debounce(memoizedFetchSteps, 500);

  useEffect(() => {
    // Initial fetch
    memoizedFetchSteps();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const updateSteps = async (newSteps: number) => {
    setSteps(newSteps);
    // Call the debounced version of fetchSteps here
    debouncedFetchSteps();
  };

  const handleStepChange = (e: React.ChangeEvent<HTMLInputElement>) => updateSteps(parseInt(e.target.value, 10));

  return (
    <div>
      <h1>Fitness Tracker</h1>
      <input
        type="number"
        placeholder="Enter Steps"
        onChange={handleStepChange} // Fixed incorrect function name
      />
      <div data-testid="real-time-data">{realTimeUpdates}</div>
      <button onClick={() => updateSteps(steps + 1000)}>Update Steps</button>
      <div data-testid="steps-display">Steps: {steps}</div> {/* Corrected closing tag */}
    </div>
  );
};

export default FitnessTracker;