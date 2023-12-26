@extends("layout", ["title" => "Home"])

@section("content")
    <body class="h-screen bg-gradient-to-r from-zinc-400 to-stone-500">

    <div class="container mx-auto mt-8 p-8 bg-white rounded shadow-md">

        <h1 class="text-2xl font-semibold mb-6">File Management</h1>

        <table class="min-w-full border border-gray-300">
            <thead>
            <tr>
                <th class="py-2 px-4 border-b">ID</th>
                <th class="py-2 px-4 border-b">Filename</th>
                <th class="py-2 px-4 border-b">Created</th>
                <th class="py-2 px-4 border-b">Accessed</th>
                <th class="py-2 px-4 border-b">Action</th>
            </tr>
            </thead>
            <tbody>
            <tr class="text-center hover:bg-gray-300 even:bg-amber-100 odd:bg-sky-100">
                <td class="py-2 px-4 border-b">1</td>
                <td class="py-2 px-4 border-b">example.txt</td>
                <td class="py-2 px-4 border-b">2023-01-01 15:00:45</td>
                <td class="py-2 px-4 border-b">2023-01-02 14:17:23</td>
                <td class="py-2 px-4 border-b">
                    <button
                        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-1 px-2 rounded-full focus:outline-none
                    focus:shadow-outline-blue active:bg-blue-800">
                        Download
                    </button>
                    <button
                        class="bg-red-500 hover:bg-red-600 text-white font-bold py-1 px-2 rounded-full focus:outline-none
                    focus:shadow-outline-red active:bg-red-800">
                        Delete
                    </button>
                </td>
            </tr>

            <tr class="text-center hover:bg-gray-300 even:bg-amber-100 odd:bg-sky-100">
                <td class="py-2 px-4 border-b">1</td>
                <td class="py-2 px-4 border-b">example.txt</td>
                <td class="py-2 px-4 border-b">2023-01-01 15:00:45</td>
                <td class="py-2 px-4 border-b">2023-01-02 14:17:23</td>
                <td class="py-2 px-4 border-b">
                    <button
                        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-1 px-2 rounded-full focus:outline-none
                    focus:shadow-outline-blue active:bg-blue-800">
                        Download
                    </button>
                    <button
                        class="bg-red-500 hover:bg-red-600 text-white font-bold py-1 px-2 rounded-full focus:outline-none
                    focus:shadow-outline-red active:bg-red-800">
                        Delete
                    </button>
                </td>
            </tr>

            </tbody>
        </table>

    </div>

    </body>

@endsection
