import React, { useEffect, useState } from 'react';
import axios from 'axios';
const apiUrl = import.meta.env.VITE_API_URL;

interface Repository {
  name: string;
  stars: number;
  forks: number;
}

interface GitHubStats {
  username: string;
  followers: number;
  following: number;
  total_stars: number;
  repositories: Repository[];
}

const App: React.FC<{ username: string }> = ({ username }) => {
  const [stats, setStats] = useState<GitHubStats | null>(null);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const response = await axios.get<GitHubStats>(`${apiUrl}/api/github-stats?username=${username}`);
        setStats(response.data);
      } catch (error) {
        console.error("Error fetching data", error);
      }
    };

    fetchStats();
  }, [username]);

  if (!stats) return <p>Loading...</p>;

  return (
    <div>
      <h1>{stats.username}</h1>
      <p>Followers: {stats.followers}</p>
      <p>Following: {stats.following}</p>
      <p>Total Stars: {stats.total_stars}</p>
      <h2>Repositories</h2>
      <ul>
        {stats.repositories.map(repo => (
          <li key={repo.name}>
            {repo.name} - ⭐ {repo.stars} | 🍴 {repo.forks}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default App;