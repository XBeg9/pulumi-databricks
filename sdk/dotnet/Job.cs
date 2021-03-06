// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks
{
    public partial class Job : Pulumi.CustomResource
    {
        [Output("createdTime")]
        public Output<int> CreatedTime { get; private set; } = null!;

        [Output("creatorUserName")]
        public Output<string> CreatorUserName { get; private set; } = null!;

        [Output("emailNotifications")]
        public Output<Outputs.JobEmailNotifications?> EmailNotifications { get; private set; } = null!;

        [Output("existingClusterId")]
        public Output<string?> ExistingClusterId { get; private set; } = null!;

        [Output("jarMainClassName")]
        public Output<string?> JarMainClassName { get; private set; } = null!;

        [Output("jarParameters")]
        public Output<ImmutableArray<string>> JarParameters { get; private set; } = null!;

        [Output("jarUri")]
        public Output<string?> JarUri { get; private set; } = null!;

        [Output("jobId")]
        public Output<int> JobId { get; private set; } = null!;

        [Output("libraryCrans")]
        public Output<ImmutableArray<Outputs.JobLibraryCran>> LibraryCrans { get; private set; } = null!;

        [Output("libraryEggs")]
        public Output<ImmutableArray<string>> LibraryEggs { get; private set; } = null!;

        [Output("libraryJars")]
        public Output<ImmutableArray<string>> LibraryJars { get; private set; } = null!;

        [Output("libraryMavens")]
        public Output<ImmutableArray<Outputs.JobLibraryMaven>> LibraryMavens { get; private set; } = null!;

        [Output("libraryPypis")]
        public Output<ImmutableArray<Outputs.JobLibraryPypi>> LibraryPypis { get; private set; } = null!;

        [Output("libraryWhls")]
        public Output<ImmutableArray<string>> LibraryWhls { get; private set; } = null!;

        [Output("maxConcurrentRuns")]
        public Output<int?> MaxConcurrentRuns { get; private set; } = null!;

        [Output("maxRetries")]
        public Output<int?> MaxRetries { get; private set; } = null!;

        [Output("minRetryIntervalMillis")]
        public Output<int?> MinRetryIntervalMillis { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("newCluster")]
        public Output<Outputs.JobNewCluster?> NewCluster { get; private set; } = null!;

        [Output("notebookBaseParameters")]
        public Output<ImmutableDictionary<string, string>?> NotebookBaseParameters { get; private set; } = null!;

        [Output("notebookPath")]
        public Output<string?> NotebookPath { get; private set; } = null!;

        [Output("pythonFile")]
        public Output<string?> PythonFile { get; private set; } = null!;

        [Output("pythonParameters")]
        public Output<ImmutableArray<string>> PythonParameters { get; private set; } = null!;

        [Output("retryOnTimeout")]
        public Output<bool?> RetryOnTimeout { get; private set; } = null!;

        [Output("schedule")]
        public Output<Outputs.JobSchedule?> Schedule { get; private set; } = null!;

        [Output("sparkSubmitParameters")]
        public Output<ImmutableArray<string>> SparkSubmitParameters { get; private set; } = null!;

        [Output("timeoutSeconds")]
        public Output<int?> TimeoutSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs? args = null, CustomResourceOptions? options = null)
            : base("databricks:index/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("databricks:index/job:Job", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
        {
            return new Job(name, id, state, options);
        }
    }

    public sealed class JobArgs : Pulumi.ResourceArgs
    {
        [Input("emailNotifications")]
        public Input<Inputs.JobEmailNotificationsArgs>? EmailNotifications { get; set; }

        [Input("existingClusterId")]
        public Input<string>? ExistingClusterId { get; set; }

        [Input("jarMainClassName")]
        public Input<string>? JarMainClassName { get; set; }

        [Input("jarParameters")]
        private InputList<string>? _jarParameters;
        public InputList<string> JarParameters
        {
            get => _jarParameters ?? (_jarParameters = new InputList<string>());
            set => _jarParameters = value;
        }

        [Input("jarUri")]
        public Input<string>? JarUri { get; set; }

        [Input("libraryCrans")]
        private InputList<Inputs.JobLibraryCranArgs>? _libraryCrans;
        public InputList<Inputs.JobLibraryCranArgs> LibraryCrans
        {
            get => _libraryCrans ?? (_libraryCrans = new InputList<Inputs.JobLibraryCranArgs>());
            set => _libraryCrans = value;
        }

        [Input("libraryEggs")]
        private InputList<string>? _libraryEggs;
        public InputList<string> LibraryEggs
        {
            get => _libraryEggs ?? (_libraryEggs = new InputList<string>());
            set => _libraryEggs = value;
        }

        [Input("libraryJars")]
        private InputList<string>? _libraryJars;
        public InputList<string> LibraryJars
        {
            get => _libraryJars ?? (_libraryJars = new InputList<string>());
            set => _libraryJars = value;
        }

        [Input("libraryMavens")]
        private InputList<Inputs.JobLibraryMavenArgs>? _libraryMavens;
        public InputList<Inputs.JobLibraryMavenArgs> LibraryMavens
        {
            get => _libraryMavens ?? (_libraryMavens = new InputList<Inputs.JobLibraryMavenArgs>());
            set => _libraryMavens = value;
        }

        [Input("libraryPypis")]
        private InputList<Inputs.JobLibraryPypiArgs>? _libraryPypis;
        public InputList<Inputs.JobLibraryPypiArgs> LibraryPypis
        {
            get => _libraryPypis ?? (_libraryPypis = new InputList<Inputs.JobLibraryPypiArgs>());
            set => _libraryPypis = value;
        }

        [Input("libraryWhls")]
        private InputList<string>? _libraryWhls;
        public InputList<string> LibraryWhls
        {
            get => _libraryWhls ?? (_libraryWhls = new InputList<string>());
            set => _libraryWhls = value;
        }

        [Input("maxConcurrentRuns")]
        public Input<int>? MaxConcurrentRuns { get; set; }

        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        [Input("minRetryIntervalMillis")]
        public Input<int>? MinRetryIntervalMillis { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("newCluster")]
        public Input<Inputs.JobNewClusterArgs>? NewCluster { get; set; }

        [Input("notebookBaseParameters")]
        private InputMap<string>? _notebookBaseParameters;
        public InputMap<string> NotebookBaseParameters
        {
            get => _notebookBaseParameters ?? (_notebookBaseParameters = new InputMap<string>());
            set => _notebookBaseParameters = value;
        }

        [Input("notebookPath")]
        public Input<string>? NotebookPath { get; set; }

        [Input("pythonFile")]
        public Input<string>? PythonFile { get; set; }

        [Input("pythonParameters")]
        private InputList<string>? _pythonParameters;
        public InputList<string> PythonParameters
        {
            get => _pythonParameters ?? (_pythonParameters = new InputList<string>());
            set => _pythonParameters = value;
        }

        [Input("retryOnTimeout")]
        public Input<bool>? RetryOnTimeout { get; set; }

        [Input("schedule")]
        public Input<Inputs.JobScheduleArgs>? Schedule { get; set; }

        [Input("sparkSubmitParameters")]
        private InputList<string>? _sparkSubmitParameters;
        public InputList<string> SparkSubmitParameters
        {
            get => _sparkSubmitParameters ?? (_sparkSubmitParameters = new InputList<string>());
            set => _sparkSubmitParameters = value;
        }

        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        public JobArgs()
        {
        }
    }

    public sealed class JobState : Pulumi.ResourceArgs
    {
        [Input("createdTime")]
        public Input<int>? CreatedTime { get; set; }

        [Input("creatorUserName")]
        public Input<string>? CreatorUserName { get; set; }

        [Input("emailNotifications")]
        public Input<Inputs.JobEmailNotificationsGetArgs>? EmailNotifications { get; set; }

        [Input("existingClusterId")]
        public Input<string>? ExistingClusterId { get; set; }

        [Input("jarMainClassName")]
        public Input<string>? JarMainClassName { get; set; }

        [Input("jarParameters")]
        private InputList<string>? _jarParameters;
        public InputList<string> JarParameters
        {
            get => _jarParameters ?? (_jarParameters = new InputList<string>());
            set => _jarParameters = value;
        }

        [Input("jarUri")]
        public Input<string>? JarUri { get; set; }

        [Input("jobId")]
        public Input<int>? JobId { get; set; }

        [Input("libraryCrans")]
        private InputList<Inputs.JobLibraryCranGetArgs>? _libraryCrans;
        public InputList<Inputs.JobLibraryCranGetArgs> LibraryCrans
        {
            get => _libraryCrans ?? (_libraryCrans = new InputList<Inputs.JobLibraryCranGetArgs>());
            set => _libraryCrans = value;
        }

        [Input("libraryEggs")]
        private InputList<string>? _libraryEggs;
        public InputList<string> LibraryEggs
        {
            get => _libraryEggs ?? (_libraryEggs = new InputList<string>());
            set => _libraryEggs = value;
        }

        [Input("libraryJars")]
        private InputList<string>? _libraryJars;
        public InputList<string> LibraryJars
        {
            get => _libraryJars ?? (_libraryJars = new InputList<string>());
            set => _libraryJars = value;
        }

        [Input("libraryMavens")]
        private InputList<Inputs.JobLibraryMavenGetArgs>? _libraryMavens;
        public InputList<Inputs.JobLibraryMavenGetArgs> LibraryMavens
        {
            get => _libraryMavens ?? (_libraryMavens = new InputList<Inputs.JobLibraryMavenGetArgs>());
            set => _libraryMavens = value;
        }

        [Input("libraryPypis")]
        private InputList<Inputs.JobLibraryPypiGetArgs>? _libraryPypis;
        public InputList<Inputs.JobLibraryPypiGetArgs> LibraryPypis
        {
            get => _libraryPypis ?? (_libraryPypis = new InputList<Inputs.JobLibraryPypiGetArgs>());
            set => _libraryPypis = value;
        }

        [Input("libraryWhls")]
        private InputList<string>? _libraryWhls;
        public InputList<string> LibraryWhls
        {
            get => _libraryWhls ?? (_libraryWhls = new InputList<string>());
            set => _libraryWhls = value;
        }

        [Input("maxConcurrentRuns")]
        public Input<int>? MaxConcurrentRuns { get; set; }

        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        [Input("minRetryIntervalMillis")]
        public Input<int>? MinRetryIntervalMillis { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("newCluster")]
        public Input<Inputs.JobNewClusterGetArgs>? NewCluster { get; set; }

        [Input("notebookBaseParameters")]
        private InputMap<string>? _notebookBaseParameters;
        public InputMap<string> NotebookBaseParameters
        {
            get => _notebookBaseParameters ?? (_notebookBaseParameters = new InputMap<string>());
            set => _notebookBaseParameters = value;
        }

        [Input("notebookPath")]
        public Input<string>? NotebookPath { get; set; }

        [Input("pythonFile")]
        public Input<string>? PythonFile { get; set; }

        [Input("pythonParameters")]
        private InputList<string>? _pythonParameters;
        public InputList<string> PythonParameters
        {
            get => _pythonParameters ?? (_pythonParameters = new InputList<string>());
            set => _pythonParameters = value;
        }

        [Input("retryOnTimeout")]
        public Input<bool>? RetryOnTimeout { get; set; }

        [Input("schedule")]
        public Input<Inputs.JobScheduleGetArgs>? Schedule { get; set; }

        [Input("sparkSubmitParameters")]
        private InputList<string>? _sparkSubmitParameters;
        public InputList<string> SparkSubmitParameters
        {
            get => _sparkSubmitParameters ?? (_sparkSubmitParameters = new InputList<string>());
            set => _sparkSubmitParameters = value;
        }

        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        public JobState()
        {
        }
    }
}
