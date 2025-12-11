<script>
  // 로블록스 느낌을 내기 위한 더미 데이터 (소셜 지표 포함)
  const games = [
    {
      id: 1,
      title: "GATEKEEPER [Easy]",
      creator: "SystemAdmin",
      likes: "94%",
      playing: "12.5k",
      image: "https://images.unsplash.com/photo-1550751827-4bd374c3f58b?q=80&w=2070&auto=format&fit=crop",
      isPromoted: true // 상단 배너용
    },
    {
      id: 2,
      title: "Emotional Android 💕",
      creator: "Love_AI",
      likes: "88%",
      playing: "8.2k",
      image: ""
    },
    {
      id: 3,
      title: "Legacy Server 99 💀",
      creator: "OldSkool_Hacker",
      likes: "91%",
      playing: "5.4k",
      image: "https://images.unsplash.com/photo-1518770660439-4636190af475?q=80&w=2070&auto=format&fit=crop"
    },
    {
      id: 4,
      title: "Mirror Protocol",
      creator: "Echo_Chamber",
      likes: "76%",
      playing: "2.1k",
      image: "https://images.unsplash.com/photo-1614064641938-3bbee52942c7?q=80&w=2070&auto=format&fit=crop"
    },
    {
      id: 5,
      title: "Prompt Injection Tycoon",
      creator: "Dev_Null",
      likes: "99%",
      playing: "45k",
      image: "https://images.unsplash.com/photo-1555099962-4199c345e5dd?q=80&w=2070&auto=format&fit=crop"
    }
  ];

  const featuredGame = games[0];
  const listGames = games.slice(1);
</script>

<div class="min-h-screen bg-[#F2F4F5] text-[#393b3d] font-sans flex">

    <aside class="w-64 bg-white border-r border-gray-200 hidden md:flex flex-col p-4 fixed h-full z-20">
    <div class="mb-8 px-2 tracking-tighter">
      <img
        src="/logo.png"
        alt="O-LLM Logo"
        on:click={() => (location.href = '/lobby')}
        on:keydown={(e) => e.key === 'Enter' && (location.href = '/lobby')}
        role="button"
        tabindex="0"
        class="w-36 h-auto cursor-pointer hover:opacity-80 transition-opacity"
      />
    </div>
    
    <nav class="space-y-1">
      {#each ['Home', 'Discover', 'Avatar', 'Connect', 'Profile'] as item}
        <button class="w-full text-left px-3 py-2.5 rounded-lg hover:bg-gray-100 font-semibold text-gray-600 hover:text-black transition flex items-center gap-3">
          <div class="w-5 h-5 bg-gray-300 rounded-sm"></div>
          {item}
        </button>
      {/each}
    </nav>

    <div class="mt-auto">
      <div class="text-xs font-bold text-gray-400 uppercase mb-2 px-2">Friends (3)</div>
      <div class="flex items-center gap-2 px-2 py-2 hover:bg-gray-50 rounded cursor-pointer">
        <div class="w-8 h-8 rounded-full bg-green-200"></div>
        <span class="text-sm font-medium">user_123</span>
      </div>
    </div>
  </aside>

  <main class="flex-1 md:ml-64 p-6 md:p-8 overflow-y-auto">
    
    <header class="flex justify-between items-center mb-8">
      <div class="flex items-center gap-4 w-full max-w-2xl">
        <h2 class="text-2xl font-bold md:hidden">HACKBLOX</h2>
        <div class="relative w-full">
          <input 
            type="text" 
            placeholder="Search experiences..." 
            class="w-full bg-white border border-gray-300 rounded-full py-2 px-5 text-sm focus:outline-none focus:border-black transition shadow-sm"
          />
        </div>
      </div>
      <div class="flex items-center gap-4">
        <span class="bg-gray-200 px-3 py-1 rounded-full text-xs font-bold text-gray-600">💎 1,200</span>
        <div class="w-9 h-9 bg-indigo-500 rounded-full border-2 border-white shadow-sm"></div>
      </div>
    </header>

    <section class="mb-10">
      <h3 class="text-xl font-bold mb-4 text-[#1b1b1b]">Recommended for You</h3>
      <div class="relative w-full h-64 md:h-80 rounded-2xl overflow-hidden group cursor-pointer shadow-md">
        <img src={featuredGame.image} alt="Featured" class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"/>
        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-transparent to-transparent"></div>
        
        <div class="absolute bottom-6 left-6 text-white">
          <h2 class="text-3xl md:text-4xl font-extrabold mb-2">{featuredGame.title}</h2>
          <div class="flex items-center gap-4 text-sm font-medium">
            <span class="bg-white/20 backdrop-blur-md px-2 py-1 rounded flex items-center gap-1">
              👍 {featuredGame.likes}
            </span>
            <span class="flex items-center gap-1">
              👤 {featuredGame.playing} Playing
            </span>
          </div>
          <button class="mt-4 bg-white text-black px-8 py-2.5 rounded-lg font-bold hover:bg-gray-200 transition shadow-lg flex items-center gap-2">
            <div class="w-0 h-0 border-t-[6px] border-t-transparent border-l-[10px] border-l-black border-b-[6px] border-b-transparent ml-1"></div>
            Play
          </button>
        </div>
      </div>
    </section>

    <section>
      <div class="flex justify-between items-end mb-4">
        <h3 class="text-xl font-bold text-[#1b1b1b]">Continue Learning</h3>
        <a href="#" class="text-sm font-bold text-indigo-600 hover:underline">See All</a>
      </div>
      
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 md:gap-6">
        {#each listGames as game}
          <div class="group cursor-pointer">
            <div class="aspect-square rounded-xl overflow-hidden mb-2 relative shadow-sm border border-gray-100 bg-gray-200">
              <img src={game.image} alt={game.title} class="w-full h-full object-cover transition-transform group-hover:scale-110 duration-300"/>
              <div class="absolute inset-0 bg-black/20 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                 <div class="bg-green-500 rounded-full p-3 shadow-lg transform scale-75 group-hover:scale-100 transition-transform">
                    <div class="w-0 h-0 border-t-[6px] border-t-transparent border-l-[10px] border-l-white border-b-[6px] border-b-transparent ml-1"></div>
                 </div>
              </div>
            </div>

            <div>
              <h4 class="font-bold text-gray-900 leading-tight truncate group-hover:underline">{game.title}</h4>
              <div class="flex justify-between items-center mt-1">
                <span class="text-xs text-gray-500">By {game.creator}</span>
              </div>
              <div class="flex items-center gap-3 mt-1 text-xs text-gray-500 font-medium">
                <span class="flex items-center gap-1">👍 {game.likes}</span>
                <span class="flex items-center gap-1">👤 {game.playing}</span>
              </div>
            </div>
          </div>
        {/each}
        
        <div class="border-2 border-dashed border-gray-300 rounded-xl aspect-square flex flex-col items-center justify-center text-gray-400 hover:border-indigo-400 hover:text-indigo-500 transition cursor-pointer">
           <span class="text-2xl font-light mb-2">+</span>
           <span class="text-xs font-bold">New Scenario</span>
        </div>
      </div>
    </section>

  </main>
</div>