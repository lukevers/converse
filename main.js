// Module to control application life.
var app = require('app');
// Module to create native browser window.
var BrowserWindow = require('browser-window');
// Module to run process
var spawn = require('child_process').spawn;

// Run process and log output
var child = spawn('./converse');
child.stdout.on('data', function(data) {
	console.log(data);
}); child.stderr.on('data', function(data) {
	console.log(data);
}); child.on('close', function(code) {
	console.log('child process exited with code ' + code);
});

// Report crashes to our server.
require('crash-reporter').start();

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the javascript object is GCed.
var mainWindow = null;

// Quit when all windows are closed.
app.on('window-all-closed', function() {
	if (process.platform != 'darwin') {
		quit();
	}
});

// This method will be called when Electron has done everything
// initialization and ready for creating browser windows.
app.on('ready', function() {
	// Create the browser window.
	mainWindow = new BrowserWindow({width: 800, height: 600});
	mainWindow.loadUrl('http://localhost:3040');

	// Emitted when the window is closed.
	mainWindow.on('closed', quit);
});

// Quit/kill everything
var quit = function() {
	app.quit();
	child.kill();
}
