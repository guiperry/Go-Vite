import React, { useState, useEffect } from 'react';
import { Play, Settings } from 'lucide-react';

function App() {
  const [status, setStatus] = useState<string>('idle');
  const [items, setItems] = useState<any[]>([]);

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const response = await fetch('/api/v1/items');
      const data = await response.json();
      setItems(data.items || []);
    } catch (error) {
      console.error('Failed to fetch items:', error);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-brand-900 to-slate-900 text-white p-6">
      <div className="max-w-7xl mx-auto">
        <div className="text-center mb-8">
          <h1 className="text-5xl font-bold mb-2 bg-gradient-to-r from-brand-400 to-blue-400 bg-clip-text text-transparent">
            example
          </h1>
          <p className="text-lg text-gray-300">A Go-Vite desktop application</p>
        </div>

        <div className="bg-slate-800/50 backdrop-blur-sm rounded-xl p-6 mb-8 border border-brand-500/30">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-2xl font-bold">Dashboard</h2>
            <span className="px-3 py-1 bg-green-600/30 text-green-300 rounded-full text-sm">
              {status}
            </span>
          </div>
          
          <div className="flex gap-3">
            <button
              onClick={() => setStatus('running')}
              className="flex items-center gap-2 bg-gradient-to-r from-brand-600 to-blue-600 px-6 py-3 rounded-lg font-semibold hover:from-brand-500 hover:to-blue-500 transition-all"
            >
              <Play className="w-5 h-5" />
              Start
            </button>
            <button className="flex items-center gap-2 bg-slate-700 px-6 py-3 rounded-lg font-semibold hover:bg-slate-600 transition-all">
              <Settings className="w-5 h-5" />
              Settings
            </button>
          </div>
        </div>

        <div className="bg-slate-800/50 backdrop-blur-sm rounded-xl p-6 border border-brand-500/30">
          <h3 className="text-xl font-bold mb-4">Items</h3>
          {items.length === 0 ? (
            <p className="text-gray-400">No items yet</p>
          ) : (
            <ul className="space-y-2">
              {items.map((item, idx) => (
                <li key={idx} className="p-3 bg-slate-700/50 rounded-lg">
                  {JSON.stringify(item)}
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
