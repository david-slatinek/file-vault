<div class="max-w-md mb-5">
    <label for="password" class="block text-gray-700 text-sm font-bold mb-2">Password</label>
    <input type="password" id="password" name="password" placeholder="Enter your password"
           class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring focus:border-blue-300 border
                           border-gray-600 border-2" minlength="12" required>
    <p class="text-xs text-gray-500 mt-1">Must be at least 12 characters, at least 1 uppercase letter,
        1 lowercase letter, 1 digit, and 1 special character.</p>
    <label for="show">Show password</label>
    <input id="show" name="show" type="checkbox" onclick="(() => {
                    let x = document.getElementById('password');
                    if (x.type === 'password') {
                        x.type = 'text';
                    } else {
                        x.type = 'password';
                    }
                })();">
</div>
