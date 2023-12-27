@extends("layout", ["title" => "2FA Setup"])

<body class="flex items-center justify-center h-screen bg-gray-100 bg-gradient-to-r from-cyan-600 to-blue-700">

@section("content")
    <div class="bg-white p-8 rounded shadow-md w-100 text-center">

        @if($otp)
            <img src="data:image/png;base64, {{$otp->url}}" alt="QR code" class="mb-4 mx-auto">
            <div>
                <h1 class="text-2xl font-semibold mb-6 text-center text-gray-800">{{$otp->key}}</h1>
            </div>
        @endif

        <div>
            <p class="text-gray-700 mb-4">Please scan the QR code with your 2FA app or enter the code manually.</p>
            <p class="text-gray-700 mb-4">As a backup, we recommend you save these codes in a password manager.</p>
        </div>

        <div class="flex justify-between items-center">
            <div></div>
            <a href="{{route("user.file")}}">
                <button
                    class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full focus:outline-none
                        focus:shadow-outline-blue active:bg-blue-800">
                    Next
                </button>
            </a>
        </div>

    </div>
@endsection

</body>

